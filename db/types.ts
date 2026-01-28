import type { ColumnType } from "kysely";
export type Generated<T> = T extends ColumnType<infer S, infer I, infer U>
  ? ColumnType<S, I | undefined, U>
  : ColumnType<T, T | undefined, T>;
export type Timestamp = ColumnType<Date, Date | string, Date | string>;

export const ScheduleStatus = {
    TODO: "TODO",
    IN_PROGRESS: "IN_PROGRESS",
    DONE: "DONE"
} as const;
export type ScheduleStatus = (typeof ScheduleStatus)[keyof typeof ScheduleStatus];
export const ActivityLogType = {
    POMODORO: "POMODORO",
    MANUAL: "MANUAL",
    STOPWATCH: "STOPWATCH",
    FREE: "FREE"
} as const;
export type ActivityLogType = (typeof ActivityLogType)[keyof typeof ActivityLogType];
export type ActivityLog = {
    id: string;
    userId: string;
    startAt: Timestamp;
    endAt: Timestamp;
    duration: number;
    type: ActivityLogType;
    scheduleId: string;
    note: string | null;
    createdAt: Generated<Timestamp>;
};
export type Event = {
    id: string;
    startAt: Timestamp;
    endAt: Timestamp;
    isAllDay: Generated<boolean>;
    recurrenceRule: string | null;
    originalEventId: string | null;
};
export type Goal = {
    id: string;
    userId: string;
    title: string;
    description: string | null;
    createdAt: Generated<Timestamp>;
    updatedAt: Timestamp;
    parentId: string | null;
};
export type Schedule = {
    id: string;
    userId: string;
    goalId: string | null;
    title: string;
    description: string | null;
    status: Generated<ScheduleStatus>;
    createdAt: Generated<Timestamp>;
    updatedAt: Timestamp;
    timezone: string;
};
export type Task = {
    id: string;
    deadlineAt: Timestamp;
    estimatedDuration: number;
    recurrenceRule: string | null;
    parentId: string | null;
};
export type User = {
    id: string;
    email: string;
    name: string | null;
    password: string;
    createdAt: Generated<Timestamp>;
    updatedAt: Timestamp;
};
export type DB = {
    ActivityLog: ActivityLog;
    Event: Event;
    Goal: Goal;
    Schedule: Schedule;
    Task: Task;
    User: User;
};

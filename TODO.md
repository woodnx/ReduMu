Follow the steps below to finish setting up your application.

## Kysely

Ensure that `DATABASE_URL` is configured as desired in `.env` file, then create the database:

```bash
pnpm kysely:migrate # creates kysely tables
```


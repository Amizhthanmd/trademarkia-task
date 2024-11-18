# Trademarkia Task: Order Inventory Management

## Getting Started

### Prerequisites

- Ensure **Go** is installed on your system.
- Ensure **PostgreSQL** is installed and running.

### Steps to Set Up the Project

1. **Clone the Repository**:

```bash
git clone https://github.com/Amizhthanmd/trademarkia-task.git
cd trademarkia-task
```

2. **Install Dependencies**:

```bash
go mod tidy
```

3. **.env file**:  
   Update the `POSTGRES_DB_URL` in the `.env` file to match your PostgreSQL setup. Below is an example:

```bash
POSTGRES_DB_URL=postgresql://postgres:your_password@127.0.0.1:5432/your_database
```

### Running Database Migrations

1. **Navigate to the migrations folder**:

```bash
cd db/migrations
```

2. **Run the migration script**:

```bash
go run main.go
```

3. **Choose one of the following options when prompted**:
   - 1: Create the database
   - 2: Run migrations
   - 3: Run triggers
   - 4: Execute all three actions

Complete the database setup based on your choice.

### Running the Application

1. **Navigate back to the root directory**:

```bash
cd ../../
```

2. **Start the application**:

```bash
go run main.go
```

### API Documentation

Refer to the [Postman API Documentation](https://martian-shadow-968002.postman.co/workspace/Amizhthan~294aad53-278f-4ed0-a7f2-383ece75cff1/collection/29108316-a6dd0b44-ba13-424e-bead-5422e8728d24?action=share&creator=29108316&active-environment=29108316-77c515df-280e-423b-ae32-9e7fc9d23ae7) for details on API usage.

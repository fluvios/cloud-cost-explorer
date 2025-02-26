# Cloud Cost Explorer - Golang Service

## Overview
The Cloud Cost Explorer Golang service fetches cloud cost data from **AWS, GCP, Azure, and Alibaba Cloud** using their respective APIs. It supports dynamic date range queries and caches results in **Redis** to improve performance.

## Features
- Fetch **cloud cost data** from AWS, GCP, Azure, and Alibaba Cloud.
- Supports **dynamic date range queries** (`YYYY-MM-DD` format).
- **Redis caching** to minimize API calls.
- **Error handling and logging** for better observability.
- Uses **environment variables** for API credentials and configuration.

## API Endpoints

### Health Check
- **Endpoint:** `GET /health`
- **Response:**
  ```json
  { "status": "ok" }
  ```

### AWS Cost Data
- **Endpoint:** `GET /costs/aws?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD`
- **Query Parameters:**
  - `start_date`: Start date (format `YYYY-MM-DD`).
  - `end_date`: End date (format `YYYY-MM-DD`).
- **Response Example:**
  ```json
  {
    "2023-06-01": "120.50",
    "2023-06-02": "115.75"
  }
  ```

### GCP Cost Data
- **Endpoint:** `GET /costs/gcp?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD`

### Azure Cost Data
- **Endpoint:** `GET /costs/azure?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD`

### Alibaba Cloud Cost Data *(Coming Soon)*
- **Endpoint:** `GET /costs/alibaba?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD`

## Environment Variables
Set the following environment variables before running the service:
```env
AWS_ACCESS_KEY_ID=your_aws_access_key
AWS_SECRET_ACCESS_KEY=your_aws_secret_key
AWS_ROLE_ARN=your_aws_role_arn
GCP_BILLING_ACCOUNT=your_gcp_billing_account
AZURE_SUBSCRIPTION_ID=your_azure_subscription_id
REDIS_HOST=localhost
REDIS_PORT=6379
```

## Setup & Run
1. **Install dependencies**
   ```sh
   go mod tidy
   ```
2. **Run the service**
   ```sh
   go run main.go
   ```
3. **Using Docker**
   ```sh
   docker build -t cloud-cost-explorer .
   docker run -p 8080:8080 --env-file .env cloud-cost-explorer
   ```

## Future Enhancements
- Implement Alibaba Cloud cost fetching.
- Add unit tests and integration testing.
- Deploy with Kubernetes.
- Implement authentication & rate limiting.

## Contributors
- **[Your Name]** - Initial development
- **[Contributors]** - Ongoing improvements

## License
This project is licensed under the **MIT License**.
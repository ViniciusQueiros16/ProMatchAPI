# ProMatch API

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go&logoColor=white)
![AWS Lambda](https://img.shields.io/badge/AWS_Lambda-Serverless-FF9900?logo=aws-lambda&logoColor=white)
![Serverless](https://img.shields.io/badge/Serverless_Framework-FD5750?logo=serverless&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-Database-4479A1?logo=mysql&logoColor=white)

**ProMatch API** is a robust, serverless backend infrastructure built in Go. It powers the ProMatch platform by providing secure, scalable endpoints for user management, service discovery, and contract negotiations between contractors and service providers.

## 🏗 Architecture

The API leverages a cloud-native **Serverless Architecture** on AWS:
- **Compute**: AWS Lambda functions written in Go for high performance and low cold start times.
- **Routing**: API Gateway routes HTTP requests to Lambda functions.
- **Database**: Relational MySQL database for structured data integrity.
- **Deployment**: Serverless Framework handles Infrastructure as Code (IaC) and deployments.

## 📋 Prerequisites

To run and deploy this project, you need:
- [Go](https://golang.org/doc/install) (Golang)
- Node.js (v20.2.0 recommended)
- [Serverless Framework](https://www.serverless.com/framework/docs/getting-started/) (`npm install -g serverless`)
- AWS CLI configured with appropriate credentials
- MySQL Server

## 🚀 Setup Instructions

1. **Database Initialization**
   In your MySQL command prompt, execute the script to create the necessary tables for ProMatch:
   ```sql
   mysql> source /path/to/create-tables.sql
   ```

2. **Environment Configuration**
   Create a `.env` file in the root directory based on `.env.example` and set your database credentials and AWS configuration:
   ```env
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_HOST=your_host
   ```

3. **Dependencies**
   Fetch Go modules:
   ```bash
   go mod download
   ```

## ☁️ Deployment

Deploy the application to AWS using the Serverless Framework:

```bash
make build
serverless deploy
```

Once deployment is successful, the terminal will output your API Gateway endpoints ready to be integrated with the ProMatch FrontEnd.

## 👨‍💻 Author

**Vinicius Queirós Muniz**
- 🔗 [GitHub Profile](https://github.com/ViniciusQueiros16)
- 💼 Software Engineer based in Salvador, BA, Brazil.

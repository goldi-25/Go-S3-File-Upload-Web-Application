# Go-S3-File-Upload-Web-Application
Descriptions:  main.go – Runs the web server, handles file uploads, connects to S3, processes requests.  index.html – The front-end page where the user selects and uploads files.  go.mod – Defines the module name and dependencies for the Go project.  go.sum – Contains checksums for dependencies (auto-generated).
# Go S3 File Upload Web App

A lightweight Go-based web application that lets users upload files directly to an Amazon S3 bucket through a simple HTML interface.  
Designed for easy deployment with full auto-deployment from GitHub using Render.

---

## 🚀 Features
- Upload files directly from the browser to S3
- Simple and clean HTML upload interface
- Go backend using AWS SDK
- Automatic deployments from GitHub to Render
- Fast, lightweight, and production-ready

---

## 📂 Project Structure
| File | Description |
|------|-------------|
| main.go | Main Go server file. Handles routes, file uploads, and S3 integration. |
| index.html | Front-end upload form displayed in the browser. |
| go.mod | Go module file defining project dependencies. |
| go.sum | Auto-generated dependency checksum file. |

---

## ⚙️ Requirements
- Go installed locally (optional for development)
- AWS S3 bucket
- AWS credentials with S3 permissions
- GitHub account
- Render account (for deployment)

---

## 🌐 Deployment (Render Auto-Deploy)
This project is deployed using Render Web Services, directly connected to GitHub.

### Steps:
1. Push your project to GitHub
2. Go to https://render.com
3. Click **New → Web Service**
4. Select your GitHub repository
5. Choose **Environment: Go**
6. Deploy the service
7. Open the service → go to **Settings**
8. Enable **Auto Deploy**
9. Add the following Environment Variables:
   - AWS_ACCESS_KEY_ID  
   - AWS_SECRET_ACCESS_KEY  
   - AWS_REGION  
   - BUCKET_NAME  

Render will now rebuild and redeploy your app automatically on every GitHub push.

---

## ▶️ Running Locally
1. Set your AWS environment variables  
2. Run the Go application  
3. Open your browser and access the HTML upload page  
4. Upload a file and it will appear in your S3 bucket  


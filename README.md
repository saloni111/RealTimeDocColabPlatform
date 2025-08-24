# 🌟 DocHub - Collaborative Writing Platform

**Built by Saloni** | *Where ideas come to life together*

---

## What is DocHub?

DocHub is a real-time collaborative document platform I built during my journey as a software developer. It's designed for teams, students, and creators who want to work together seamlessly on documents - whether you're brainstorming project ideas, writing research papers, or collaborating on creative content.

## Why I Built This

As someone passionate about building meaningful software, I wanted to create something that solved a real problem. Having worked on group projects and collaborative writing, I experienced firsthand how frustrating it can be when multiple people try to edit the same document. DocHub was born from that frustration - a platform where collaboration feels natural and ideas flow freely.

## ✨ What Makes DocHub Special

### 🚀 **Real-Time Magic**
- Multiple people can edit the same document simultaneously
- See changes as they happen - no refresh needed
- Track who made what changes with smart versioning

### 🎯 **User-Friendly Design**
- Clean, intuitive interface that feels familiar
- No complicated setup - just sign up and start writing
- Works beautifully on desktop and mobile

### 🔧 **Built for Developers**
- Microservices architecture for scalability
- RESTful APIs that are easy to integrate
- Modern tech stack with Go, gRPC, and DynamoDB

### 🤝 **Team-Focused Features**
- User management with secure authentication
- Document sharing with granular permissions
- Real-time synchronization that just works

## 🛠️ How It's Built

I chose to build DocHub using modern, scalable technologies:

**Backend Services:**
- **Go** - For high-performance microservices
- **gRPC** - Lightning-fast inter-service communication
- **DynamoDB** - NoSQL database for flexible document storage
- **JWT** - Secure user authentication

**Architecture:**
- **API Gateway** - Single entry point for all requests
- **User Service** - Handles authentication and user management
- **Document Service** - Manages document CRUD operations
- **Collaboration Service** - Powers real-time editing features

## 🚀 Getting Started

### Prerequisites
- Go 1.22 or higher
- Docker (for DynamoDB Local)

### Quick Setup

1. **Clone the repository**
```bash
   git clone https://github.com/yourusername/collaborative-doc-platform
   cd collaborative-doc-platform
```

2. **Start DynamoDB Local**
```bash
   docker run -p 9000:8000 amazon/dynamodb-local
```

3. **Create required tables**
```bash
   # User table
   aws dynamodb create-table --table-name doc_users \
     --attribute-definitions AttributeName=user_id,AttributeType=S \
     --key-schema AttributeName=user_id,KeyType=HASH \
  --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
  --endpoint-url http://localhost:9000

   # Documents table
   aws dynamodb create-table --table-name docs \
     --attribute-definitions AttributeName=document_id,AttributeType=S \
     --key-schema AttributeName=document_id,KeyType=HASH \
  --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
  --endpoint-url http://localhost:9000
```

4. **Start all services**
```bash
./start_all.sh
```

5. **Open DocHub**
   Navigate to `http://localhost:8080` and start collaborating!

## 🎮 How to Use DocHub

### Getting Started
1. **Sign Up** - Create your account in the "Get Started" section
2. **Create a Document** - Use "Create Something Amazing" to write your first doc
3. **Share & Collaborate** - Give the document ID to teammates using "Team Up"
4. **Edit Together** - Use "Share Your Thoughts" for real-time editing

### Pro Tips
- **Document IDs** are automatically generated when you create documents
- **Version History** tracks all changes with timestamps
- **Real-time Sync** happens automatically when multiple people edit

## 🏗️ Project Structure

```
📁 collaborative-doc-platform/
├── 🌐 api-gateway/          # HTTP API and web interface
├── 👤 user-service/         # User authentication & management
├── 📄 document-service/     # Document CRUD operations
├── 🤝 collaboration-service/ # Real-time editing features
├── 📋 proto/               # Protocol buffer definitions
└── 🚀 start_all.sh        # Quick startup script
```

## 🔧 API Reference

### User Management
- `POST /user/register` - Create new user account
- `POST /login` - Authenticate user
- `GET /user?user_id=X` - Get user profile

### Document Operations
- `POST /document/create` - Create new document
- `GET /document/{id}` - Retrieve document
- `PUT /document/{id}` - Update document content
- `DELETE /document/{id}` - Delete document

### Collaboration Features
- `POST /document/join/{id}` - Join document for editing
- `POST /document/sync/{id}` - Sync real-time changes
- `POST /document/leave/{id}` - Leave collaborative session

## 🎯 Technical Achievements

This project demonstrates several advanced concepts:

- **Microservices Architecture** with proper service separation
- **gRPC Communication** for high-performance inter-service calls
- **Real-time Collaboration** using WebSocket-like protocols
- **JWT Authentication** with secure token management
- **NoSQL Database Design** with proper indexing
- **API Gateway Pattern** for service orchestration
- **Document Versioning** with timestamp tracking

## 🌟 What I Learned

Building DocHub taught me:
- How to design scalable microservices architectures
- The importance of real-time data synchronization
- Best practices for API design and documentation
- How to balance performance with user experience
- The value of testing and iterative development

## 🚀 Future Enhancements

Ideas I'm excited to implement:
- **Rich Text Editor** with formatting options
- **Document Templates** for common use cases
- **Comment System** for better collaboration
- **File Upload** for images and attachments
- **Export Options** (PDF, Word, etc.)
- **Mobile App** for on-the-go editing

## 🤝 Contributing

This is a portfolio project, but I'm always open to feedback and suggestions! If you have ideas for improvements or find any issues, feel free to:
- Open an issue
- Submit a pull request
- Reach out to me directly

## 📧 Connect With Me

I'd love to hear your thoughts about DocHub or discuss opportunities!
- **LinkedIn**: [Your LinkedIn Profile]
- **Email**: [your.email@example.com]
- **Portfolio**: [Your Portfolio Website]

## 📝 License

This project is built for educational and portfolio purposes. Feel free to explore the code and learn from it!

---

**Built with ❤️ by Saloni** | *Making collaboration effortless, one document at a time*
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/saloni111/RealTimeDocColabPlatform/api-gateway/handler"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Add actual health checks for dependent services
	health := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"version":   "1.0.0",
		"services": map[string]string{
			"user-service":          "unknown", // TODO: Check service health
			"document-service":      "unknown",
			"collaboration-service": "unknown",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Return JSON if requested
	if r.Header.Get("Accept") == "application/json" {
		response := map[string]interface{}{
			"message": "DocHub - Collaborative Document Platform",
			"status":  "running",
			"endpoints": map[string]string{
				"user_register":   "POST /user/register",
				"user_login":      "POST /login",
				"user_profile":    "GET /user",
				"create_document": "POST /document/create",
				"get_document":    "GET /document/{document_id}",
				"update_document": "PUT /document/{document_id}",
				"delete_document": "DELETE /document/{document_id}",
				"list_documents":  "GET /documents",
				"join_document":   "POST /document/join/{document_id}",
				"sync_document":   "POST /document/sync/{document_id}",
				"leave_document":  "POST /document/leave/{document_id}",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Cache-busting headers
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "text/html")

	// Google Docs inspired interface
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>DocHub</title>
    <link href="https://fonts.googleapis.com/css2?family=Google+Sans:wght@300;400;500;600&display=swap" rel="stylesheet">
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: 'Google Sans', -apple-system, BlinkMacSystemFont, sans-serif;
            background-color: #f9fbfd;
            color: #202124;
            line-height: 1.6;
        }
        
        /* Google Docs Header */
        .header {
            background: #fff;
            border-bottom: 1px solid #e8eaed;
            padding: 0 24px;
            height: 64px;
            display: flex;
            align-items: center;
            justify-content: space-between;
            box-shadow: 0 1px 3px rgba(60,64,67,.3);
        }
        
        .logo {
            display: flex;
            align-items: center;
            font-size: 22px;
            font-weight: 500;
            color: #5f6368;
        }
        
        .logo .material-icons {
            color: #4285f4;
            margin-right: 8px;
            font-size: 24px;
        }
        
        .user-info {
            display: flex;
            align-items: center;
            gap: 16px;
            color: #5f6368;
            font-size: 14px;
        }
        
        /* Main Container */
        .container {
            max-width: 1200px;
            margin: 0 auto;
            padding: 32px 24px;
        }
        
        /* Document List */
        .documents-section {
            background: #fff;
            border-radius: 8px;
            box-shadow: 0 1px 3px rgba(60,64,67,.3);
            margin-bottom: 24px;
        }
        
        .section-header {
            padding: 24px 24px 16px;
            border-bottom: 1px solid #e8eaed;
            display: flex;
            align-items: center;
            justify-content: space-between;
        }
        
        .section-title {
            font-size: 16px;
            font-weight: 500;
            color: #202124;
            display: flex;
            align-items: center;
        }
        
        .section-title .material-icons {
            margin-right: 8px;
            color: #5f6368;
        }
        
        /* Google Material Button */
        .btn {
            background: #4285f4;
            color: #fff;
            border: none;
            border-radius: 4px;
            padding: 8px 16px;
            font-size: 14px;
            font-weight: 500;
            cursor: pointer;
            display: inline-flex;
            align-items: center;
            gap: 8px;
            transition: all 0.2s;
            text-decoration: none;
        }
        
        .btn:hover {
            background: #3367d6;
            box-shadow: 0 1px 3px rgba(60,64,67,.3);
        }
        
        .btn-outline {
            background: transparent;
            color: #4285f4;
            border: 1px solid #dadce0;
        }
        
        .btn-outline:hover {
            background: #f8f9fa;
        }
        
        /* Document Grid */
        .documents-grid {
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
            gap: 16px;
            padding: 24px;
        }
        
        .document-card {
            background: #fff;
            border: 1px solid #e8eaed;
            border-radius: 8px;
            padding: 16px;
            cursor: pointer;
            transition: all 0.2s;
            position: relative;
        }
        
        .document-actions {
            position: absolute;
            top: 8px;
            right: 8px;
            display: flex;
            gap: 4px;
            opacity: 0.8;
            transition: opacity 0.2s;
            z-index: 100;
            background: rgba(255,255,255,0.9);
            border-radius: 4px;
            padding: 2px;
        }
        
        .document-card:hover .document-actions {
            opacity: 1;
            background: rgba(255,255,255,1);
        }
        
        .action-btn {
            background: rgba(60,64,67,0.1);
            border: none;
            border-radius: 50%;
            width: 32px;
            height: 32px;
            display: flex;
            align-items: center;
            justify-content: center;
            cursor: pointer;
            transition: all 0.2s;
        }
        
        .action-btn:hover {
            background: rgba(60,64,67,0.2);
        }
        
        .action-btn.delete:hover {
            background: rgba(234,67,53,0.1);
            color: #ea4335;
        }
        
        .action-btn.share:hover {
            background: rgba(66,133,244,0.1);
            color: #4285f4;
        }
        
        .document-card:hover {
            box-shadow: 0 2px 8px rgba(60,64,67,.3);
            border-color: #4285f4;
        }
        
        .document-preview {
            background: #f8f9fa;
            border-radius: 4px;
            height: 120px;
            margin-bottom: 12px;
            padding: 12px;
            font-size: 12px;
            color: #5f6368;
            overflow: hidden;
            position: relative;
        }
        
        .document-preview::after {
            content: '';
            position: absolute;
            bottom: 0;
            left: 0;
            right: 0;
            height: 20px;
            background: linear-gradient(transparent, #f8f9fa);
        }
        
        .document-title {
            font-size: 14px;
            font-weight: 500;
            color: #202124;
            margin-bottom: 4px;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
        }
        
        .document-meta {
            font-size: 12px;
            color: #5f6368;
            display: flex;
            align-items: center;
            gap: 8px;
        }
        
        .document-meta .material-icons {
            font-size: 16px;
        }
        
        /* Create New Document */
        .create-new {
            background: #fff;
            border: 2px dashed #dadce0;
            border-radius: 8px;
            padding: 32px;
            text-align: center;
            cursor: pointer;
            transition: all 0.2s;
            color: #5f6368;
        }
        
        .create-new:hover {
            border-color: #4285f4;
            background: #f8f9fa;
        }
        
        .create-new .material-icons {
            font-size: 48px;
            color: #4285f4;
            margin-bottom: 8px;
        }
        
        /* Modal */
        .modal {
            display: none;
            position: fixed;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
            background: rgba(60,64,67,.3);
            z-index: 1000;
            align-items: center;
            justify-content: center;
        }
        
        .modal.show {
            display: flex;
        }
        
        .modal-content {
            background: #fff;
            border-radius: 8px;
            box-shadow: 0 8px 32px rgba(60,64,67,.3);
            max-width: 500px;
            width: 90%;
            padding: 24px;
        }
        
        .modal-header {
            display: flex;
            align-items: center;
            justify-content: space-between;
            margin-bottom: 24px;
        }
        
        .modal-title {
            font-size: 20px;
            font-weight: 500;
            color: #202124;
        }
        
        .close-btn {
            background: none;
            border: none;
            font-size: 24px;
            color: #5f6368;
            cursor: pointer;
            padding: 8px;
            border-radius: 50%;
        }
        
        .close-btn:hover {
            background: #f8f9fa;
        }
        
        /* Form Styles */
        .form-group {
            margin-bottom: 16px;
        }
        
        .form-label {
            display: block;
            font-size: 14px;
            font-weight: 500;
            color: #202124;
            margin-bottom: 8px;
        }
        
        .form-input {
            width: 100%;
            padding: 12px 16px;
            border: 1px solid #dadce0;
            border-radius: 4px;
            font-size: 14px;
            transition: all 0.2s;
        }
        
        .form-input:focus {
            outline: none;
            border-color: #4285f4;
            box-shadow: 0 0 0 2px rgba(66,133,244,.2);
        }
        
        .form-textarea {
            min-height: 120px;
            resize: vertical;
            font-family: inherit;
        }
        
        /* Editor */
        .editor-container {
            background: #fff;
            border-radius: 8px;
            box-shadow: 0 1px 3px rgba(60,64,67,.3);
            margin-bottom: 24px;
            min-height: 500px;
            display: none;
        }
        
        .editor-header {
            padding: 16px 24px;
            border-bottom: 1px solid #e8eaed;
            display: flex;
            align-items: center;
            justify-content: space-between;
        }
        
        .editor-title {
            font-size: 18px;
            font-weight: 500;
            color: #202124;
        }
        
        .editor-actions {
            display: flex;
            gap: 8px;
        }
        
        .editor-content {
            padding: 24px;
            min-height: 400px;
            font-size: 14px;
            line-height: 1.6;
            outline: none;
        }
        
        /* Status indicators */
        .status {
            display: inline-flex;
            align-items: center;
            gap: 4px;
            font-size: 12px;
            color: #5f6368;
        }
        
        .status.online {
            color: #34a853;
        }
        
        .status .material-icons {
            font-size: 16px;
        }
        
        /* Loading */
        .loading {
            display: inline-flex;
            align-items: center;
            gap: 8px;
            color: #5f6368;
            font-size: 14px;
        }
        
        .spinner {
            width: 16px;
            height: 16px;
            border: 2px solid #e8eaed;
            border-top: 2px solid #4285f4;
            border-radius: 50%;
            animation: spin 1s linear infinite;
        }
        
        @keyframes spin {
            0% { transform: rotate(0deg); }
            100% { transform: rotate(360deg); }
        }
        
        /* Empty state */
        .empty-state {
            text-align: center;
            padding: 48px 24px;
            color: #5f6368;
        }
        
        .empty-state .material-icons {
            font-size: 64px;
            color: #dadce0;
            margin-bottom: 16px;
        }
        
        .empty-state h3 {
            font-size: 18px;
            font-weight: 400;
            margin-bottom: 8px;
        }
        
        .empty-state p {
            font-size: 14px;
        }
    </style>
</head>
<body>
    <!-- Header -->
    <div class="header">
        <div class="logo">
            <span class="material-icons">description</span>
            DocHub
        </div>
        <div class="user-info">
            <span class="status online">
                <span class="material-icons">circle</span>
                Online
            </span>
            <span id="currentUser">Guest User</span>
        </div>
    </div>

    <!-- Main Container -->
    <div class="container">
        <!-- Documents Section -->
        <div class="documents-section">
            <div class="section-header">
                <div class="section-title">
                    <span class="material-icons">folder</span>
                    Recent documents
                </div>
                <button class="btn" onclick="showCreateModal()">
                    <span class="material-icons">add</span>
                    Create new
                </button>
            </div>
            
            <div class="documents-grid" id="documentsGrid">
                <div class="create-new" onclick="showCreateModal()">
                    <div class="material-icons">add</div>
                    <div>Create a new document</div>
                </div>
            </div>
        </div>

        <!-- Document Editor (Hidden by default) -->
        <div class="editor-container" id="editorContainer">
            <div class="editor-header">
                <div class="editor-title" id="editorTitle">Untitled Document</div>
                <div class="editor-actions">
                    <span class="status" id="syncStatus" title="Document sync status">
                        <span class="material-icons">check_circle</span>
                        <span>Saved</span>
                    </span>
                    <button class="btn-outline" onclick="shareDocument()" id="shareBtn">
                        <span class="material-icons">share</span>
                        Share
                    </button>
                    <button class="btn-outline" onclick="closeEditor()">
                        <span class="material-icons">close</span>
                        Close
                    </button>
                </div>
            </div>
            <div class="editor-content" id="editorContent" contenteditable="true">
                Start typing your document...
            </div>
        </div>
    </div>

    <!-- Create Document Modal -->
    <div class="modal" id="createModal">
        <div class="modal-content">
            <div class="modal-header">
                <div class="modal-title">Create new document</div>
                <button class="close-btn" onclick="hideCreateModal()">
                    <span class="material-icons">close</span>
                </button>
            </div>
            
            <form onsubmit="createNewDocument(event)">
                <div class="form-group">
                    <label class="form-label">Document title</label>
                    <input type="text" class="form-input" id="newDocTitle" placeholder="Untitled document" required>
                </div>
                
                <div class="form-group">
                    <label class="form-label">Your name</label>
                    <input type="text" class="form-input" id="authorName" placeholder="Your name" required>
                </div>
                
                <div class="form-group">
                    <label class="form-label">Initial content (optional)</label>
                    <textarea class="form-input form-textarea" id="initialContent" placeholder="Start writing..."></textarea>
                </div>
                
                <div style="display: flex; gap: 8px; justify-content: flex-end; margin-top: 24px;">
                    <button type="button" class="btn-outline" onclick="hideCreateModal()">Cancel</button>
                    <button type="submit" class="btn">
                        <span class="material-icons">add</span>
                        Create
                    </button>
                </div>
            </form>
        </div>
    </div>

    <script>
        // Global state
        let currentUser = 'Guest User';
        let documents = JSON.parse(localStorage.getItem('dochub_documents') || '[]');
        let currentDocument = null;
        let saveTimeout = null;

        // Initialize
        document.addEventListener('DOMContentLoaded', function() {
            loadDocumentsFromServer();
            setupAutoSave();
            checkForSharedDocument();
        });

        // Modal functions
        function showCreateModal() {
            document.getElementById('createModal').classList.add('show');
            document.getElementById('newDocTitle').focus();
        }

        function hideCreateModal() {
            document.getElementById('createModal').classList.remove('show');
            document.getElementById('newDocTitle').value = '';
            document.getElementById('authorName').value = '';
            document.getElementById('initialContent').value = '';
        }

        // Create new document
        async function createNewDocument(event) {
            event.preventDefault();
            
            const title = document.getElementById('newDocTitle').value;
            const author = document.getElementById('authorName').value;
            const content = document.getElementById('initialContent').value || 'Start writing...';
            
            try {
                showSyncStatus('Creating...', 'loading');
                
                const response = await fetch('/document/create', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ title, content, author })
                });
                
                const result = await response.json();
                
                if (result.document_id) {
                    const doc = {
                        id: result.document_id,
                        title,
                        author,
                        content,
                        created: new Date().toISOString(),
                        modified: new Date().toISOString()
                    };
                    
                    documents.unshift(doc);
                    localStorage.setItem('dochub_documents', JSON.stringify(documents));
                    
                    hideCreateModal();
                    // Refresh from server to ensure we have the latest data
                    loadDocumentsFromServer().then(() => {
                        openDocument(doc);
                    });
                    showSyncStatus('Document created', 'success');
                } else {
                    throw new Error('Failed to create document');
                }
            } catch (error) {
                console.error('Error creating document:', error);
                showSyncStatus('Error creating document', 'error');
            }
        }

        // Load documents from server first, then localStorage as fallback
        async function loadDocumentsFromServer() {
            try {
                showSyncStatus('Loading documents...', 'loading');
                const response = await fetch('/documents');
                if (response.ok) {
                    const data = await response.json();
                    if (data.documents && data.documents.length >= 0) { // Changed > 0 to >= 0 to handle empty list
                        documents = data.documents;
                        // Also update localStorage with fresh data
                        localStorage.setItem('dochub_documents', JSON.stringify(documents));
                        showSyncStatus('Documents loaded (' + documents.length + ')', 'success');
                        loadDocuments();
                        return Promise.resolve();
                    }
                }
            } catch (error) {
                console.error('Error loading from server:', error);
            }
            
            // Fallback to localStorage
            const storedDocs = JSON.parse(localStorage.getItem('dochub_documents') || '[]');
            documents = storedDocs;
            showSyncStatus('Loaded from cache', 'success');
            loadDocuments();
            return Promise.resolve();
        }

        // Load documents
        function loadDocuments() {
            const grid = document.getElementById('documentsGrid');
            const createButton = grid.querySelector('.create-new');
            
            // Clear existing documents but keep create button
            grid.innerHTML = '';
            grid.appendChild(createButton);
            
            if (documents.length === 0) {
                const emptyState = document.createElement('div');
                emptyState.className = 'empty-state';
                emptyState.innerHTML = 
                    '<div class="material-icons">description</div>' +
                    '<h3>No documents yet</h3>' +
                    '<p>Create your first document to get started</p>';
                grid.appendChild(emptyState);
                return;
            }
            
            documents.forEach(doc => {
                const card = document.createElement('div');
                card.className = 'document-card';
                card.onclick = (e) => {
                    if (!e.target.closest('.document-actions')) {
                        openDocument(doc);
                    }
                };
                
                const preview = doc.content.substring(0, 150) + (doc.content.length > 150 ? '...' : '');
                const modifiedDate = new Date(doc.modified).toLocaleDateString();
                
                card.innerHTML = 
                    '<div class="document-actions">' +
                        '<button class="action-btn share" onclick="shareDocumentFromCard(\'' + doc.id + '\', \'' + doc.title + '\', event)" title="Share document">' +
                            '<span class="material-icons">share</span>' +
                        '</button>' +
                        '<button class="action-btn delete" onclick="deleteDocumentFromCard(\'' + doc.id + '\', event)" title="Delete document">' +
                            '<span class="material-icons">delete</span>' +
                        '</button>' +
                    '</div>' +
                    '<div class="document-preview">' + preview + '</div>' +
                    '<div class="document-title">' + doc.title + '</div>' +
                    '<div class="document-meta">' +
                        '<span class="material-icons">person</span>' +
                        '<span>' + doc.author + '</span>' +
                        '<span>•</span>' +
                        '<span>' + modifiedDate + '</span>' +
                    '</div>';
                
                grid.appendChild(card);
            });
        }

        // Open document in editor
        async function openDocument(doc) {
            currentDocument = doc;
            
            // Show loading
            showSyncStatus('Loading...', 'loading');
            
            try {
                // Always fetch fresh content from server
                const response = await fetch('/document/' + doc.id);
                if (response.ok) {
                    const freshDoc = await response.json();
                    
                    // Update current document with fresh data
                    currentDocument.content = freshDoc.content;
                    currentDocument.title = freshDoc.title;
                    
                    // Update in local storage too
                    const docIndex = documents.findIndex(d => d.id === doc.id);
                    if (docIndex !== -1) {
                        documents[docIndex].content = freshDoc.content;
                        documents[docIndex].title = freshDoc.title;
                        localStorage.setItem('dochub_documents', JSON.stringify(documents));
                    }
                    
                    document.getElementById('editorTitle').textContent = freshDoc.title;
                    // Set content as plain text to avoid HTML issues
                    const editorContent = document.getElementById('editorContent');
                    editorContent.innerHTML = ''; // Clear first
                    editorContent.innerText = freshDoc.content; // Then set clean text
                    document.getElementById('editorContainer').style.display = 'block';
                    
                    showSyncStatus('Loaded', 'success');
                    
                    // Hide documents list
                    document.querySelector('.documents-section').style.display = 'none';
                    
                    // Start collaboration sync
                    startCollaborationSync();
                } else {
                    throw new Error('Failed to load document');
                }
            } catch (error) {
                console.error('Error loading document:', error);
                showSyncStatus('Error loading', 'error');
                
                // Fallback to local content
                document.getElementById('editorTitle').textContent = doc.title;
                document.getElementById('editorContent').innerText = doc.content;
                document.getElementById('editorContainer').style.display = 'block';
                document.querySelector('.documents-section').style.display = 'none';
                
                // Start collaboration sync even with fallback
                startCollaborationSync();
            }
        }

        // Close editor
        function closeEditor() {
            if (currentDocument) {
                saveDocument();
            }
            
            // Stop collaboration sync
            stopCollaborationSync();
            
            document.getElementById('editorContainer').style.display = 'none';
            document.querySelector('.documents-section').style.display = 'block';
            currentDocument = null;
        }

        // Auto-save setup
        function setupAutoSave() {
            const editor = document.getElementById('editorContent');
            
            // Track ALL editing activity to prevent sync interruptions
            editor.addEventListener('input', function() {
                if (currentDocument) {
                    lastTypingTime = Date.now();
                    showSyncStatus('Saving...', 'loading');
                    
                    clearTimeout(saveTimeout);
                    saveTimeout = setTimeout(() => {
                        saveDocument();
                    }, 1000);
                }
            });
            
            // Track ALL possible user interactions with the editor
            ['keydown', 'keyup', 'keypress', 'input', 'paste', 'cut', 'focus', 'click', 'mousedown', 'mouseup', 'touchstart', 'touchend'].forEach(eventType => {
                editor.addEventListener(eventType, function() {
                    lastTypingTime = Date.now();
                    console.log('🖱️ User activity detected:', eventType);
                });
            });
        }

        // Save document
        async function saveDocument() {
            if (!currentDocument) return;
            
            // Get both HTML and text content
            const editorElement = document.getElementById('editorContent');
            const htmlContent = editorElement.innerHTML;
            const textContent = editorElement.innerText || editorElement.textContent;
            
            // Use text content for saving to prevent HTML pollution
            const contentToSave = textContent.trim() || htmlContent;
            
            try {
                const response = await fetch('/document/' + currentDocument.id, {
                    method: 'PUT',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ content: contentToSave })
                });
                
                if (response.ok) {
                    // Immediately fetch fresh content from server to verify save
                    const verifyResponse = await fetch('/document/' + currentDocument.id);
                    if (verifyResponse.ok) {
                        const freshDoc = await verifyResponse.json();
                        
                        // Update editor with the actual saved content
                        const editorElement = document.getElementById('editorContent');
                        editorElement.innerHTML = ''; // Clear first
                        editorElement.innerText = freshDoc.content; // Set fresh content
                        
                        // Update current document with fresh data
                        currentDocument.content = freshDoc.content;
                        currentDocument.title = freshDoc.title;
                        currentDocument.modified = new Date().toISOString();
                        
                        // Update localStorage with fresh data
                        const docIndex = documents.findIndex(d => d.id === currentDocument.id);
                        if (docIndex !== -1) {
                            documents[docIndex] = currentDocument;
                            localStorage.setItem('dochub_documents', JSON.stringify(documents));
                        }
                        
                        showSyncStatus('💾 Saved & verified', 'success');
                        console.log('✅ Save verified - content:', freshDoc.content.substring(0, 50) + '...');
                        
                        // Refresh document grid to show updated previews immediately
                        loadDocumentsFromServer();
                    } else {
                        showSyncStatus('💾 Saved (verification failed)', 'success');
                    }
                     
                     // Force immediate sync check to reflect changes
                     setTimeout(() => {
                         if (refreshInterval) {
                             // Trigger an immediate sync check
                             const syncEvent = new Event('sync-check');
                             document.dispatchEvent(syncEvent);
                         }
                     }, 500);
                } else {
                    throw new Error('Save failed');
                }
            } catch (error) {
                console.error('Error saving document:', error);
                showSyncStatus('Error saving', 'error');
            }
        }

        // Share document from card
        function shareDocumentFromCard(docId, docTitle, event) {
            event.stopPropagation();
            
            const shareUrl = window.location.origin + '?doc=' + docId;
            
            // Try to copy to clipboard first
            navigator.clipboard.writeText(shareUrl).then(() => {
                // Show success message
                const shareBtn = event.target.closest('.action-btn');
                const originalHTML = shareBtn.innerHTML;
                shareBtn.innerHTML = '<span class="material-icons">check</span>';
                shareBtn.style.background = 'rgba(52,168,83,0.1)';
                shareBtn.style.color = '#34a853';
                
                // Show alert with the share link
                alert('✅ Share link copied to clipboard!\n\n' +
                      'Document: "' + docTitle + '"\n' +
                      'Link: ' + shareUrl + '\n\n' +
                      'Anyone with this link can view and edit this document.\n\n' +
                      'To collaborate:\n' +
                      '1. Send this link to others\n' +
                      '2. They can open it in their browser\n' +
                      '3. Changes will sync automatically!');
                
                // Reset button after 3 seconds
                setTimeout(() => {
                    shareBtn.innerHTML = originalHTML;
                    shareBtn.style.background = '';
                    shareBtn.style.color = '';
                }, 3000);
            }).catch(() => {
                // Fallback: show the link in an alert
                alert('📋 Share this document:\n\n' +
                      'Document: "' + docTitle + '"\n' +
                      'Link: ' + shareUrl + '\n\n' +
                      'Copy this link and share it with others to collaborate!');
            });
        }

        // Delete document from card
        async function deleteDocumentFromCard(docId, event) {
            event.stopPropagation();
            
            const doc = documents.find(d => d.id === docId);
            if (!doc) return;
            
            if (!confirm('Delete "' + doc.title + '"? This action cannot be undone.')) {
                return;
            }
            
            try {
                const response = await fetch('/document/' + docId, {
                    method: 'DELETE'
                });
                
                if (response.ok) {
                    // Remove from local storage
                    documents = documents.filter(d => d.id !== docId);
                    localStorage.setItem('dochub_documents', JSON.stringify(documents));
                    
                    // Refresh document grid from server to ensure sync
                    loadDocumentsFromServer();
                    
                    // Close editor if this document is open
                    if (currentDocument && currentDocument.id === docId) {
                        closeEditor();
                    }
                } else {
                    throw new Error('Delete failed');
                }
            } catch (error) {
                console.error('Error deleting document:', error);
                alert('Failed to delete document. Please try again.');
            }
        }

        // Show sync status
        function showSyncStatus(message, type) {
            const status = document.getElementById('syncStatus');
            const icon = status.querySelector('.material-icons');
            const textSpan = status.querySelector('span:last-child');
            
            status.className = 'status';
            
            switch (type) {
                case 'loading':
                    icon.textContent = 'sync';
                    status.classList.add('loading');
                    status.style.color = '#5f6368';
                    break;
                case 'success':
                    icon.textContent = 'check_circle';
                    status.classList.add('online');
                    status.style.color = '#34a853';
                    break;
                case 'error':
                    icon.textContent = 'error';
                    status.style.color = '#ea4335';
                    break;
            }
            
            textSpan.textContent = message;
            
            if (type === 'success' || type === 'error') {
                setTimeout(() => {
                    status.className = 'status';
                    status.style.color = '#34a853';
                    icon.textContent = 'check_circle';
                    textSpan.textContent = 'Saved';
                }, 2000);
            }
        }

        // Share document
        function shareDocument() {
            if (!currentDocument) return;
            
            const shareUrl = window.location.origin + '?doc=' + currentDocument.id;
            
            // Try to copy to clipboard first
            navigator.clipboard.writeText(shareUrl).then(() => {
                const shareBtn = document.getElementById('shareBtn');
                const originalText = shareBtn.innerHTML;
                shareBtn.innerHTML = '<span class="material-icons">check</span>Copied!';
                
                // Show detailed sharing information
                alert('✅ Share link copied to clipboard!\n\n' +
                      'Document: "' + currentDocument.title + '"\n' +
                      'Link: ' + shareUrl + '\n\n' +
                      '🎯 How to collaborate:\n\n' +
                      '1. Send this link to others via email, chat, or any messaging app\n' +
                      '2. They click the link and the document opens automatically\n' +
                      '3. Everyone can edit simultaneously - changes sync every 2 seconds\n' +
                      '4. No accounts needed - just the link!\n\n' +
                      '💡 Pro tip: Open the link in multiple browser tabs to test collaboration!');
                
                setTimeout(() => {
                    shareBtn.innerHTML = originalText;
                }, 3000);
            }).catch(() => {
                // Fallback: show the link in an alert
                alert('📋 Share this document:\n\n' +
                      'Document: "' + currentDocument.title + '"\n' +
                      'Link: ' + shareUrl + '\n\n' +
                      'Copy this link and share it with others to collaborate!\n\n' +
                      '🎯 How it works:\n' +
                      '• Send the link to others\n' +
                      '• They open it and can edit immediately\n' +
                      '• Changes sync automatically every 2 seconds');
            });
        }

        // Auto-refresh document content every 5 seconds when editing
        let refreshInterval = null;
        let lastTypingTime = 0;
        
        function startCollaborationSync() {
            if (refreshInterval) clearInterval(refreshInterval);
            
            refreshInterval = setInterval(async () => {
                if (currentDocument && document.getElementById('editorContainer').style.display === 'block') {
                    try {
                        const response = await fetch('/document/' + currentDocument.id);
                        if (response.ok) {
                            const freshDoc = await response.json();
                            const editor = document.getElementById('editorContent');
                            const currentEditorContent = editor.innerText || editor.textContent || '';
                            
                            // ABSOLUTE protection - NO sync updates when editor has ANY focus or recent activity
                            const editorHasFocus = document.activeElement === editor;
                            const isRecentlyActive = (Date.now() - lastTypingTime) < 10000; // 10 seconds buffer
                            const shouldBlockSync = editorHasFocus || isRecentlyActive;
                            
                            // Debug logging for sync decisions
                            if (shouldBlockSync) {
                                console.log('🚫 Sync BLOCKED - User active (focus:', editorHasFocus, 'recent activity:', isRecentlyActive, ')');
                            } else if (freshDoc.content === currentEditorContent) {
                                console.log('✅ Sync skipped - Content identical');
                            } else if (freshDoc.content === currentDocument.content) {
                                console.log('✅ Sync skipped - No new changes');
                            } else if (currentEditorContent.trim() === '') {
                                console.log('✅ Sync skipped - Editor empty');
                            }
                            
                            // ONLY update if user is completely away from editor AND content actually changed
                            if (!shouldBlockSync && 
                                freshDoc.content !== currentEditorContent && 
                                freshDoc.content !== currentDocument.content &&
                                currentEditorContent.trim() !== '') {
                                
                                console.log('📝 Collaboration sync: SAFE UPDATE (user completely inactive)');
                                
                                // Store current scroll position
                                const scrollTop = editor.scrollTop;
                                
                                // Update content without touching cursor
                                editor.textContent = freshDoc.content;
                                
                                // Restore scroll position
                                editor.scrollTop = scrollTop;
                                
                                // Update current document
                                currentDocument.content = freshDoc.content;
                                currentDocument.title = freshDoc.title || currentDocument.title;
                                
                                // Update localStorage
                                const docIndex = documents.findIndex(d => d.id === currentDocument.id);
                                if (docIndex !== -1) {
                                    documents[docIndex] = currentDocument;
                                    localStorage.setItem('dochub_documents', JSON.stringify(documents));
                                }
                                
                                showSyncStatus('✨ Updated by collaborator', 'success');
                                
                                // Also refresh document grid to show updated previews
                                loadDocumentsFromServer();
                            } else if (freshDoc.content === currentEditorContent) {
                                // In perfect sync - show this occasionally
                                if (Math.random() < 0.1) { // 10% chance to show "in sync"
                                    showSyncStatus('🔄 In sync', 'success');
                                }
                            }
                        }
                    } catch (error) {
                        console.error('Sync error:', error);
                        showSyncStatus('⚠️ Sync error', 'error');
                    }
                }
            }, 2000); // Check every 2 seconds for better collaboration
        }
        
        function stopCollaborationSync() {
            if (refreshInterval) {
                clearInterval(refreshInterval);
                refreshInterval = null;
            }
        }

        // Check if opening shared document
        function checkForSharedDocument() {
            const urlParams = new URLSearchParams(window.location.search);
            const docId = urlParams.get('doc');
            
            if (docId) {
                // Try to load the shared document
                fetch('/document/' + docId)
                    .then(response => response.json())
                    .then(doc => {
                        // Add to local documents if not already there
                        if (!documents.find(d => d.id === docId)) {
                            const newDoc = {
                                id: doc.document_id,
                                title: doc.title,
                                author: doc.author,
                                content: doc.content,
                                created: doc.versions[0] || new Date().toISOString(),
                                modified: doc.versions[doc.versions.length - 1] || new Date().toISOString()
                            };
                            documents.unshift(newDoc);
                            localStorage.setItem('dochub_documents', JSON.stringify(documents));
                        }
                        
                        // Open the document
                        const docToOpen = documents.find(d => d.id === docId) || {
                            id: doc.document_id,
                            title: doc.title,
                            author: doc.author,
                            content: doc.content,
                            created: new Date().toISOString(),
                            modified: new Date().toISOString()
                        };
                        
                        openDocument(docToOpen);
                        loadDocuments(); // Refresh the grid
                    })
                    .catch(error => {
                        console.error('Error loading shared document:', error);
                        alert('Could not load shared document. Please check the link.');
                    });
            }
        }

        // Keyboard shortcuts
        document.addEventListener('keydown', function(e) {
            if (e.ctrlKey || e.metaKey) {
                switch (e.key) {
                    case 's':
                        e.preventDefault();
                        if (currentDocument) saveDocument();
                        break;
                    case 'n':
                        e.preventDefault();
                        showCreateModal();
                        break;
                }
            }
        });
    </script>
</body>
</html>`

	w.Write([]byte(html))
}

func main() {
	r := mux.NewRouter()

	// Root route
	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/health", healthHandler).Methods("GET")

	// User routes
	r.HandleFunc("/user/register", handler.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/login", handler.LoginUserHandler).Methods("POST")
	r.HandleFunc("/user", handler.GetUserProfileHandler).Methods("GET")

	// Document routes
	r.HandleFunc("/documents", handler.ListDocumentsHandler).Methods("GET")
	r.HandleFunc("/document/create", handler.CreateDocumentHandler).Methods("POST")
	r.HandleFunc("/document/{document_id}", handler.GetDocumentHandler).Methods("GET")
	r.HandleFunc("/document/{document_id}", handler.DeleteDocumentHandler).Methods("DELETE")
	r.HandleFunc("/document/{document_id}", handler.UpdateDocumentHandler).Methods("PUT")
	r.HandleFunc("/document/{document_id}/version", handler.ListDocumentVersionHandler).Methods("GET")

	// Collaboration routes
	r.HandleFunc("/document/join/{document_id}", handler.JoinDocumentHandler).Methods("POST")
	r.HandleFunc("/document/sync/{document_id}", handler.SyncChangesHandler).Methods("POST")
	r.HandleFunc("/document/leave/{document_id}", handler.LeaveDocumentHandler).Methods("POST")

	log.Printf("🚀 DocHub API Gateway listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

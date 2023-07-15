run:
	go run cmd/main.go

# Deploy to Google Cloud Run
# Path: Makefile
deploy:
	gcloud app deploy

# Show in browser
browser:
	gcloud app browse

logs:
	gcloud app logs tail -s default

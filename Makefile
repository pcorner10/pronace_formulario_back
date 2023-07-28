run:
	go run main.go

# Deploy to Google Cloud Run
# Path: Makefile
deploy:
	gcloud app deploy

# Show in browser
browser:
	gcloud app browse

logs:
	gcloud app logs tail -s default

rundeploy:
	gcloud run deploy quickstart-instance \
	--region=us-central1 \
	--source=. \
	--set-env-vars INSTANCE_CONNECTION_NAME="${GOOGLE_CLOUD_PROJECT}:us-central1:quickstart-instance" \
	--set-env-vars DB_NAME="quickstart_db" \
	--set-env-vars DB_USER="quickstart-service-account@${GOOGLE_CLOUD_PROJECT}.iam" \
	--service-account="quickstart-service-account@${GOOGLE_CLOUD_PROJECT}.iam.gserviceaccount.com" \
	--allow-unauthenticated

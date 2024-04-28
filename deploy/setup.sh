docker build -t ykhdr/construction-report ../construction-organization-report
docker build -t ykhdr/construction-system ../construction-organization-system
docker build -t ykhdr/construction-front ../construction-organization-front

helm upgrade --install construction ../.helm \
  --set images.repository.username=USERNAME \
  --set images.repository.password=PASSWORD \
  --set secrets.database.user=DBUSER \
  --set secrets.database.password=DBPASSWORD \
  --set secrets.database.host=DBHOST \
  --set secrets.database.port=DBPORT \
  --set secrets.database.name=DBNAME
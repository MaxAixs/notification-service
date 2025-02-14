# Delete notifications older than 5 days
psql -U postgres -d notifications -c "DELETE FROM notifications WHERE created_at < NOW() - INTERVAL '5 days';"



# API
## Input
```json
{
    "SQL": "CREATE TABLE ..."
}
```

## Output

```json
{
  "Code": 999,
  "Result": "type TutorialsTbl struct {\n\tTutorialId     int          `json:\"tutorial_id\" xorm:\"not null pk autoincr INT(11) 'tutorial_id'\"`\n\tTutorialTitle  string       `json:\"tutorial_title\" xorm:\"not null VARCHAR(100) 'tutorial_title'\"`\n\tTutorialAuthor string       `json:\"tutorial_author\" xorm:\"not null VARCHAR(40) 'tutorial_author'\"`\n\tSubmissionDate sql.NullTime `json:\"submission_date\" xorm:\"DATE 'submission_date'\"`\n}\n\n"
}
```
> Code: 999/444 (Success / Error)
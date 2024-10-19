const smtpHost = process.env.SMTP_SERVER;
const smtpPort = process.env.SMTP_PORT;
const smtpEmail = process.env.SMTP_USERNAME;
const smtpPassword = process.env.SMTP_PASSWORD;


db = connect('mongodb://mongo:27017/sinergiaManager');

const exist = db.configs.findOne({})

if (!exist) { // this part shouldnt exist in production

  db.configs.insertOne({
    SmtpHost: smtpHost,
    SmtpPort: parseInt(smtpPort, 10),
    SmtpEmail: smtpEmail,
    SmtpPassword: smtpPassword
  });
}
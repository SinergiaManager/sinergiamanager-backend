// this part shouldnt exist in production

const smtpHost = process.env.SMTP_SERVER;
const smtpPort = process.env.SMTP_PORT;
const smtpEmail = process.env.SMTP_USERNAME;
const smtpPassword = process.env.SMTP_PASSWORD;
const supportEmail = process.env.SUPPORT_EMAIL;
const mongoUri = process.env.MONGO_URI || 'mongodb://mongo:27017';
const mongoDb = process.env.MONGO_DB || 'sinergiaManager';

db = connect(mongoUri + '/' + mongoDb);

const exist = db.configs.findOne({})

console.log('Config exist: ', exist);

if (!exist) {
  db.configs.insertOne({
    smtp_host: smtpHost,
    smtp_port: parseInt(smtpPort, 10),
    smtp_user: smtpEmail,
    smtp_pass: smtpPassword,
    support_email: supportEmail
  });
}
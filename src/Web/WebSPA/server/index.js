import express from 'express';
import path from 'path';
import API from '../src/api';
import serverRenderer from './middleware/renderer';

const PORT = 5000;
const shortBaseURL = `http://localhost:5000`;

const app = express();
const router = express.Router();
const currentURL = req => `${req.protocol}://${req.get('host')}${req.originalUrl}`;

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'pug');

router.use('^/$', serverRenderer);

// other static resources should just be served as they are
router.use(express.static(path.resolve(__dirname, '..', 'build'), { maxAge: '30d' }));

// tell the app to use the above rules
app.use(router);

app.get('/:sid', async (req, res, next) => {
  const { sid } = req.params;
  const url = currentURL(req);

  API.get(`url?shortUrl=${shortBaseURL}/${sid}`)
    .then(response => {
      // Success
      const { data } = response.data;
      res.redirect(301, data.long_url);
    })
    .catch(error => {
      // Error
      if (error.response) {
        const { status } = error.response.data;
        if (status === 404) {
          res.status(404).render('404', { url, title: '404 | Page Not Found' });
        }
      } else if (error.code === 'ECONNABORTED') {
        // The request was made but no response was received
        console.log('TIMEOUT: ', error.message);
      } else {
        // Something happened in setting up the request that triggered an Error
        console.log('Request', error.request);
      }
    });
});

app.get('*', (req, res) => {
  res.status(404).render('404', { url: currentURL(req), title: '404 | Page Not Found' });
});

// start the app
app.listen(PORT, error => {
  if (error) {
    return console.log('something bad happened', error);
  }

  console.log(`listening on ${PORT}...`);
});

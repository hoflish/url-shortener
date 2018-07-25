import express from 'express';
import path from 'path';
import API from '../src/api';
import serverRenderer from './middleware/renderer';

const PORT = 5000;
const shortBaseURL = `http://localhost:5000`;
const sidRegex = /^([a-zA-Z0-9_-]){9,12}$/; // carachters that make short id (sid)

const app = express();
const currentURL = req => `${req.protocol}://${req.get('host')}${req.originalUrl}`;

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'pug');

app.use(express.static(path.resolve(__dirname, '..', 'media')));

// hande server-side redirection for valid short URL
// and diplay server errors if any
app.get('/:sid', async (req, res, next) => {
  const url = currentURL(req);
  const { sid } = req.params;

  if (!sidRegex.test(sid)) {
    // Reduce requests to service by validating sid param
    res.status(404).render('error', {
      message: `404 Not Found - the page ${url} does not exist.`,
    });
  } else {
    API.get(`url?shortUrl=${shortBaseURL}/${sid}`)
      .then(response => {
        // Success
        const { data } = response.data;
        res.redirect(301, data.long_url);
      })
      .catch(error => {
        // Error
        if (error.response) {
          const { status, message } = error.response.data;
          if (status === 404) {
            res.status(status).render('error', {
              message: `404 Not Found - the page ${url} does not exist.`,
            });

          } else {
            res.status(status).render('error', {
              message,
            });
          }
        } else {
          res.status(500).render('error', {
            message: 'Oops! something went wrong, Please try again in sometime.',
          });
        }
      });
  }
});

const router = express.Router();

// Here, we render our React SPA that creates short URL
// and notify for client errors
router.use('^/$', serverRenderer);

// other static resources should just be served as they are
router.use(express.static(path.resolve(__dirname, '..', 'build')));

// tell the app to use the above rules
app.use(router);

app.get('*', (req, res) => {
  res.status(404).render('error', {
    message: `The page ${currentURL(req)} does not exist.`,
  });
});

// start the app
app.listen(PORT, error => {
  if (error) {
    return console.log('something bad happened', error);
  }
  console.log(`listening on ${PORT}...`);
});

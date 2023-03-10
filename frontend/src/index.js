import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import reportWebVitals from './reportWebVitals';
import AOS from 'aos';
import 'aos/dist/aos.css';
import { ChakraProvider, extendTheme } from '@chakra-ui/react';
import { BrowserRouter } from 'react-router-dom';
import { Provider } from 'react-redux';
import store from './redux/store';

AOS.init({ duration: 1000 });

const colors = {
  accentLight: {
    900: '#1A365D',
    800: '#2A4365',
    700: '#2C5282',
    600: '#2B6CB0',
    500: '#3182CE',
    400: '#4299E1',
    300: '#63B3ED',
    200: '#90CDF4',
    100: '#BEE3F8',
    50: '#EBF8FF',
  },
  accentDark: {
    900: '#5F370E',
    800: '#744210',
    700: '#975A16',
    600: '#B7791F',
    500: '#D69E2E',
    400: '#ECC94B',
    300: '#F6E05E',
    200: '#FAF089',
    100: '#FEFCBF',
    50: '#FFFFF0',
  },
};

const theme = extendTheme({ colors });

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <Provider store={store}>
      <BrowserRouter>
        <ChakraProvider theme={theme}>
          <App />
        </ChakraProvider>
      </BrowserRouter>
    </Provider>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();

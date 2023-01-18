export const BASE_URL = '' + process.env.REACT_APP_LEITBE_API_URL;

//USER AUTH
export const ENDPOINT_API_POST_LOGIN = `${BASE_URL}/login`;
export const ENDPOINT_API_POST_REGISTER = `${BASE_URL}/register`;
export const ENDPOINT_API_POST_USER = `${BASE_URL}/user/`;

//MODULE ARTICLE
export const ENDPOINT_API_GET_INFORMASI = `${BASE_URL}/informasi`;
export const ENDPOINT_API_GET_FAQ = `${BASE_URL}/faq`;
export const ENDPOINT_API_GET_MODUL = `${BASE_URL}/modul`;
export const ENDPOINT_API_GET_MODUL_DETAIL = `${BASE_URL}/detail`;

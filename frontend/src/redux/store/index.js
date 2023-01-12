import { configureStore } from '@reduxjs/toolkit';
import authSliceRegister from '../slice/register';

const store = configureStore({
  reducer: {
    register: authSliceRegister,
  },
});

export default store;

import create from 'zustand';
import { devtools, persist } from 'zustand/middleware';

const useLoginState = create(
  persist(
    devtools((set) => ({
      isLoggedIn: false,
      userId: '',
      dataId: '',
      message: '',
      token: '',
      setIsLoggedIn: () => set({ isLoggedIn: true }),
      setIsLoggedOut: () => set({ isLoggedIn: false }),
      setUserId: (id) => set({ userId: id }),
      setDataId: (dataId) => set({ dataId: dataId }),
      setMessage: (message) => set({ message: message }),
      setToken: (token) => set({ token: token }),
    })),
    {
      name: 'client',
      getStorage: () => localStorage,
    }
  )
);

export default useLoginState;

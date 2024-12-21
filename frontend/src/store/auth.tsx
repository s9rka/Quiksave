import { atomWithStorage } from "jotai/utils";
import { getDefaultStore } from "jotai/vanilla";

export const authTokenAtom = atomWithStorage<string | null>(
  "authToken",
  null,
  {
    getItem: (key, initialValue) => {
      const value = localStorage.getItem(key);
      return value !== null ? JSON.parse(value) : initialValue;
    },
    setItem: (key, value) => {
      localStorage.setItem(key, JSON.stringify(value));
    },
    removeItem: (key) => {
      localStorage.removeItem(key);
    },
  },
  {
    getOnInit: true, // in an SPA with getOnInit either not set or false you will always get the initial value instead of the stored value on initialization
  }
);

export const usernameAtom = atomWithStorage<string | null>("username", null);

export const getToken = () => getDefaultStore().get(authTokenAtom);

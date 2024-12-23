import { atom, getDefaultStore } from "jotai";
import { atomWithStorage } from "jotai/utils";

export interface AuthState {
  token: string | null;
  username: string | null;
}

const initialAuthState: AuthState = {
  token: null,
  username: null,
};

export const authAtom = atomWithStorage<AuthState>("authState", initialAuthState);

export const setAuthState = (authState: AuthState) => {
  const store = getDefaultStore();
  store.set(authAtom, authState);
};

export const getAuthState = (): AuthState => {
  const store = getDefaultStore();
  return store.get(authAtom);
};

export const isAuthenticatedAtom = atom((get) => {
  const { token, username } = get(authAtom);
  return !!token && !!username;
});

export const getToken = (): string | null => {
  const { token } = getAuthState();
  return token && token.trim() !== "" ? token : null;
};

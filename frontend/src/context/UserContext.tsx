
import { useState, useEffect } from "react";
import { atom, useAtom } from "jotai";
import { atomWithStorage } from 'jotai/utils'

import { getUser } from "@/services/api";

interface UserAtomProps {
  username: string,
  email: string,

}
export const userAtom = atomWithStorage<UserAtomProps | null>("userAtom", null);

export function useInitializeAuth() {
  const [user, setUser] = useAtom(userAtom);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const verifyUser = async () => {
      try {
        const response = await getUser();
        if (response?.data) {
          console.log("User authenticated:", response.data);
          setUser(response.data);
        } else {
          console.log("No user session found");
          setUser(null);
        }
      } catch (error) {
        console.error("Error verifying user:", error);
        setUser(null);
      } finally {
        setLoading(false);
      }
    };

    verifyUser();
  }, [setUser]);

  return loading;
}
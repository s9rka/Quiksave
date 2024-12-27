// src/state/userAtom.ts
// src/hooks/useInitializeAuth.ts
import { useState, useEffect } from "react";
import { atom, useAtom } from "jotai";
import { getUser } from "@/services/api";

interface UserAtomProps {
  username: string,
  email: string,

}
export const userAtom = atom<UserAtomProps | null>(null);


export function useInitializeAuth() {
  const [user, setUser] = useAtom(userAtom);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const verifyUser = async () => {
      try {
        const response = await getUser(); // e.g. GET /api/me
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
import { useLogout } from '@/services/mutations';
import { useEffect } from 'react';

const Logout = () => {
    const logoutMutation = useLogout()

    useEffect(() => {
        const handleLogout = async () => {
          try {
            await logoutMutation.mutateAsync(); // Using mutateAsync for cleaner async handling
          } catch (error) {
            console.error('Logout failed:', error);
          }
        };
    
        handleLogout();
      }, [logoutMutation]);

  return null
}

export default Logout

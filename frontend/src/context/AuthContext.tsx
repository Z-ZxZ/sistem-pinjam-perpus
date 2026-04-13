'use client';

import React, { createContext, useContext, useEffect, useState } from 'react';

interface User {
  id: number;
  name: string;
  email: string;
  role: string;
}

interface AuthContextType {
  user: User | null;
  token: string | null;
  isLoggedIn: boolean;
  isAdmin: boolean;
  isLoading: boolean;
  isRestored: boolean; // Flag penanda kalo kita udah kelar ngecek localStorage nih
  login: (token: string, user: User) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [token, setToken] = useState<string | null>(null);
  const [isRestored, setIsRestored] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const savedToken = localStorage.getItem('token');
    const savedUser = localStorage.getItem('user');

    console.log('[AuthContext] Initializing hydration. Found token:', !!savedToken);

    if (savedToken && savedUser) {
      try {
        const parsedUser = JSON.parse(savedUser);
        console.log('[AuthContext] Hydrated user:', parsedUser.email, 'Role:', parsedUser.role);
        // eslint-disable-next-line react-hooks/set-state-in-effect
        setToken(savedToken);
        setUser(parsedUser);
      } catch (e) {
        console.error('[AuthContext] Failed to parse saved user', e);
        localStorage.removeItem('token');
        localStorage.removeItem('user');
      }
    }
    
    setIsRestored(true);
    setIsLoading(false);
  }, []);

  const login = (newToken: string, newUser: User) => {
    console.log('[AuthContext] Logging in user:', newUser.email, 'Role:', newUser.role);
    localStorage.setItem('token', newToken);
    localStorage.setItem('user', JSON.stringify(newUser));
    setToken(newToken);
    setUser(newUser);
  };

  const logout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    setToken(null);
    setUser(null);
  };

  return (
    <AuthContext.Provider
      value={{
        user,
        token,
        isLoggedIn: !!token,
        isAdmin: user?.role === 'admin',
        isLoading,
        isRestored,
        login,
        logout,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}

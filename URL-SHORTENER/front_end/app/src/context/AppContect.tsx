import { createContext, useState } from "react";

export const AppContext = createContext();

export function AppProvider({ children }) {
  const [count, setCount] = useState(0);

  return (
    <AppContext.Provider value={{ count, setCount }}>
      {children}
    </AppContext.Provider>
  );
}

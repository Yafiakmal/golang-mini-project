import { createContext, useContext, useState, type ReactNode } from "react";

const GlobalContext = createContext<any>(null);

export const GlobalProvider = ({ children }: { children: ReactNode }) => {
    const [toggle, setToggle] = useState(false);
    return (
        <GlobalContext.Provider value={{ toggle, setToggle }}>
            {children}
        </GlobalContext.Provider>
    );
};

export const useGlobal = () => useContext(GlobalContext);
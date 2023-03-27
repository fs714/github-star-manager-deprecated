import React, { createContext, ReactElement, useContext, useState } from 'react';

export interface ITheme {
  menuStyle?: 'dark' | 'light';
}

export const ThemeConfig: ITheme = {
  menuStyle: 'light'
};

const ThemeContext = createContext(
  {} as {
    theme: ITheme;
    setTheme: React.Dispatch<React.SetStateAction<ITheme>>;
  }
);

export const ThemeContextProvider = (props: { children: ReactElement }) => {
  const [theme, setTheme] = useState<ITheme>(ThemeConfig);
  return <ThemeContext.Provider value={{ theme, setTheme }}>{props.children}</ThemeContext.Provider>;
};

export const useThemeContext = () => {
  return useContext(ThemeContext);
};

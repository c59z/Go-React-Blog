import type { ReactNode } from "react";
import { RouterProvider } from "react-router-dom";
import { router } from "@/router/routes.tsx";
import { ThemeProvider } from "@emotion/react";
import theme from "@/theme";
import { CssBaseline } from "@mui/material";

interface Props {
  children?: ReactNode;
}

const AppProviders = ({ children }: Props) => {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline></CssBaseline>
      <RouterProvider router={router} />
      {children}
    </ThemeProvider>
  );
};

export default AppProviders;

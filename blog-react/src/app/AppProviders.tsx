import type { ReactNode } from "react";
import { RouterProvider } from "react-router-dom";
import { router } from "@/router/routes.tsx";
import { ThemeProvider } from "@emotion/react";
import theme from "@/theme";
import { CssBaseline } from "@mui/material";
import { SnackbarProvider } from "notistack";

interface Props {
  children?: ReactNode;
}

const AppProviders = ({ children }: Props) => {
  return (
    <ThemeProvider theme={theme}>
      <SnackbarProvider
        maxSnack={3}
        autoHideDuration={3000}
        anchorOrigin={{
          vertical: "top",
          horizontal: "center",
        }}
      >
        <CssBaseline></CssBaseline>
        <RouterProvider router={router} />
        {children}
      </SnackbarProvider>
    </ThemeProvider>
  );
};

export default AppProviders;

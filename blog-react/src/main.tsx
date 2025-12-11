import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import { RouterProvider } from "react-router-dom";
import { router } from "./router/routes.tsx";
import { setGlobalTitle } from "./utils/websiteUtils.ts";
import App from "./App.tsx";

setGlobalTitle();
// todo icon

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <RouterProvider router={router} />
    <App></App>
  </StrictMode>
);

import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import { setGlobalTitle } from "./utils/websiteUtils.ts";
import App from "./app/App.tsx";
import AppProviders from "./app/AppProviders.tsx";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import "dayjs/locale/zh-cn";

setGlobalTitle();
dayjs.extend(relativeTime);
dayjs.locale("zh-cn");
// todo icon

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <AppProviders>
      <App />
    </AppProviders>
  </StrictMode>
);

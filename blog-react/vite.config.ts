import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import path from "node:path";
import { env } from "node:process";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "src"),
    },
  },
  server: {
    host: "0.0.0.0",
    // port: 80,
    proxy: {
      "/api": {
        target: "http://127.0.0.1:8080",
        changeOrigin: true,
      },
      "/upload": {
        target: env.VITE_SERVER_URL,
        changeOrigin: true,
      },
    },
  },
});

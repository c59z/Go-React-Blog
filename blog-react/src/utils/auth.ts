import { useUserStore } from "@/stores/user";

const TOKEN_KEY = "access_token";
const EXPIRES_KEY = "access_token_expires_at";

export const auth = {
  setAuth(token: string, expiresAt: number) {
    localStorage.setItem(TOKEN_KEY, token);
    localStorage.setItem(EXPIRES_KEY, String(expiresAt));
  },

  getToken(): string | null {
    return localStorage.getItem(TOKEN_KEY);
  },

  isExpired(): boolean {
    const expiresAt = Number(localStorage.getItem(EXPIRES_KEY));
    if (!expiresAt) return true;

    const exp = expiresAt * 1000;

    return Date.now() >= exp;
  },

  isLogin(): boolean {
    return !!this.getToken() && !this.isExpired();
  },

  clear() {
    localStorage.removeItem(TOKEN_KEY);
    localStorage.removeItem(EXPIRES_KEY);
  },

  logout() {
    auth.clear();
    useUserStore.getState().clearUser();
  },
};

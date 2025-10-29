import type { User } from "@/api/user";
import { create } from "zustand";
import { persist } from "zustand/middleware";

interface UserState {
  user: User | null;
  setUser: (u: User) => void;
  updateUser: (partial: Partial<User>) => void;
  clearUser: () => void;
}

export const useUserStore = create<UserState, [["zustand/persist", UserState]]>(
  persist(
    (set) => ({
      user: null,

      setUser: (u) => set({ user: u }),

      updateUser: (partial) =>
        set((state) => ({
          user: state.user ? { ...state.user, ...partial } : null,
        })),

      clearUser: () => set({ user: null }),
    }),
    {
      name: "user-storage",
    }
  )
);

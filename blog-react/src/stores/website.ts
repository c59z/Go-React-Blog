import { create } from "zustand";
import { websiteDefaultConfig } from "@/components/config/websiteDefault";
import { isEmpty } from "@/utils/stringUtils";

export type WebsiteConfig = typeof websiteDefaultConfig;

interface WebsiteState {
  website: WebsiteConfig;
  setWebsite: (data: Partial<WebsiteConfig>) => void;
  resetWebsite: () => void;
}

export const useWebsiteStore = create<WebsiteState>((set) => ({
  website: websiteDefaultConfig,

  setWebsite: (data) =>
    set(() => ({
      website: {
        ...websiteDefaultConfig,
        ...Object.fromEntries(
          Object.entries(data).filter(([, value]) => !isEmpty(value))
        ),
      },
    })),

  resetWebsite: () =>
    set(() => ({
      website: websiteDefaultConfig,
    })),
}));

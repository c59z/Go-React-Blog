// src/app/AppInit.tsx
import { useEffect } from "react";
import { useWebsiteStore } from "@/stores/website";
import { websiteInfo } from "@/api/website";

const AppInit = () => {
  const setWebsite = useWebsiteStore((s) => s.setWebsite);

  useEffect(() => {
    websiteInfo().then((res) => {
      setWebsite(res.data);
    });
  }, [setWebsite]);

  return null;
};

export default AppInit;

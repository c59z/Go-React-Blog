import { Typography } from "@mui/material";
import clsx from "clsx";
import defaultLogo from "@/assets/coffee.png";
import "./index.scss";
import { useWebsiteStore } from "@/stores/website";

interface LogoProps {
  size?: "sm" | "md";
}

const Logo = ({ size = "md" }: LogoProps) => {
  const { full_logo, title, slogan } = useWebsiteStore((s) => s.website);

  return (
    <div className={clsx("logo-root-style", `logo-${size}-style`)}>
      <img
        src={full_logo || defaultLogo}
        alt="logo"
        className="logo-image-style"
      />

      <div className="logo-text-wrapper-style">
        <Typography className="logo-title-style">{title}</Typography>
        <Typography className="logo-subtitle-style">{slogan}</Typography>
      </div>
    </div>
  );
};

export default Logo;

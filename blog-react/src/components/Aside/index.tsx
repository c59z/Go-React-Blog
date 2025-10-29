import type { ReactNode } from "react";
import "./index.scss";

interface Props {
  children: ReactNode;
}

const Aside = ({ children }: Props) => {
  return <aside className="aside">{children}</aside>;
};

export default Aside;

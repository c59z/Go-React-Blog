import type { HeadingItem } from "../../ArticlePage";
import "./index.scss";
import { useScrollSpyWithClickLock } from "./hook/useScrollSpyWithClickLock";

interface Props {
  headings: HeadingItem[];
}

const ArticleTOC = ({ headings }: Props) => {
  const { activeId, scrollTo } = useScrollSpyWithClickLock(headings, {
    offset: 96,
    lockDuration: 400,
  });

  const aside = headings && headings.length && (
    <aside className="article-toc">
      <div className="toc-title">Index</div>

      <ul>
        {headings.map((h) => (
          <li
            key={h.id}
            className={`level-${h.level} ${activeId === h.id ? "active" : ""}`}
            onClick={() => scrollTo(h.id)}
          >
            {h.text}
          </li>
        ))}
      </ul>
    </aside>
  );

  return aside || null;
};
export default ArticleTOC;

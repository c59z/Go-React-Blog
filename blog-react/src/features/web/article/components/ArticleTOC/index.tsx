import { useEffect, useState } from "react";
import type { HeadingItem } from "../../ArticlePage";

interface Props {
  headings: HeadingItem[];
}

const ArticleTOC = ({ headings }: Props) => {
  const [activeId, setActiveId] = useState("");

  useEffect(() => {
    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((e) => {
          if (e.isIntersecting) {
            setActiveId(e.target.id);
          }
        });
      },
      { rootMargin: "0px 0px -70% 0px" }
    );

    headings.forEach((h) => {
      const el = document.getElementById(h.id);
      if (el) observer.observe(el);
    });

    return () => observer.disconnect();
  }, [headings]);

  const scrollTo = (id: string) => {
    document.getElementById(id)?.scrollIntoView({
      behavior: "smooth",
    });
  };

  return (
    <aside className="article-toc">
      <div className="toc-title">目录</div>

      <ul>
        {headings.map((h) => (
          <li
            key={h.id}
            className={`${activeId === h.id ? "active" : ""} level-${h.level}`}
            onClick={() => scrollTo(h.id)}
          >
            {h.text}
          </li>
        ))}
      </ul>
    </aside>
  );
};

export default ArticleTOC;

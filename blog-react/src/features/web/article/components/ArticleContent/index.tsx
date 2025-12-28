import { useEffect, useRef } from "react";
import type { HeadingItem } from "../../ArticlePage";

interface Props {
  onHeadingsChange: (list: HeadingItem[]) => void;
}

const ArticleContent = ({ onHeadingsChange }: Props) => {
  const contentRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (!contentRef.current) return;

    const nodes = contentRef.current.querySelectorAll("h2, h3");

    const list: HeadingItem[] = Array.from(nodes).map((node) => {
      const id =
        (node.id ||
          node.textContent?.trim().replace(/\s+/g, "-").toLowerCase()) ??
        "";

      node.id = id;

      return {
        id,
        text: node.textContent ?? "",
        level: Number(node.tagName.replace("H", "")),
      };
    });

    onHeadingsChange(list);
  }, [onHeadingsChange]);

  return <div ref={contentRef}>{"文章内容"}</div>;
};

export default ArticleContent;

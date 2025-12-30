import { useEffect, useRef } from "react";
import type { HeadingItem } from "../../ArticlePage";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import rehypeSlug from "rehype-slug";
import rehypeAutolinkHeadings from "rehype-autolink-headings";
import "./index.scss";

interface Props {
  onHeadingsChange: (list: HeadingItem[]) => void;
  content: string;
}

const ArticleContent = ({ onHeadingsChange, content }: Props) => {
  const contentRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (!contentRef.current) return;

    const nodes = contentRef.current.querySelectorAll("h2, h3");

    const list: HeadingItem[] = Array.from(nodes).map((node) => ({
      id: node.id,
      text: node.textContent ?? "",
      level: Number(node.tagName.replace("H", "")),
    }));

    onHeadingsChange(list);
  }, [content, onHeadingsChange]);

  return (
    <div ref={contentRef} className="article-content">
      <ReactMarkdown
        remarkPlugins={[remarkGfm]}
        rehypePlugins={[rehypeSlug, rehypeAutolinkHeadings]}
      >
        {content}
      </ReactMarkdown>
    </div>
  );
};

export default ArticleContent;

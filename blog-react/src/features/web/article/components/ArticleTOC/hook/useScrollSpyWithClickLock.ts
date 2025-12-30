import { useEffect, useRef, useState } from "react";
import type { HeadingItem } from "../../../ArticlePage";

interface Options {
  offset?: number;
  lockDuration?: number;
}

export const useScrollSpyWithClickLock = (
  headings: HeadingItem[],
  options: Options = {}
) => {
  const { offset = 96, lockDuration = 400 } = options;

  const [activeId, setActiveId] = useState("");
  const isClickScrollingRef = useRef(false);
  const lockTimerRef = useRef<number | null>(null);

  const scrollTo = (id: string) => {
    isClickScrollingRef.current = true;
    setActiveId(id);

    document.getElementById(id)?.scrollIntoView({
      behavior: "smooth",
      block: "start",
    });

    if (lockTimerRef.current) {
      window.clearTimeout(lockTimerRef.current);
    }

    lockTimerRef.current = window.setTimeout(() => {
      isClickScrollingRef.current = false;
    }, lockDuration);
  };

  useEffect(() => {
    if (!headings.length) return;

    let ticking = false;

    const onScroll = () => {
      if (ticking || isClickScrollingRef.current) return;

      ticking = true;
      requestAnimationFrame(() => {
        let nextId: string | null = null;

        for (const h of headings) {
          const el = document.getElementById(h.id);
          if (!el) continue;

          const top = el.getBoundingClientRect().top - offset;

          if (top <= 0) {
            nextId = h.id;
          } else {
            break;
          }
        }

        if (nextId && nextId !== activeId) {
          setActiveId(nextId);
        }

        ticking = false;
      });
    };

    window.addEventListener("scroll", onScroll, { passive: true });
    onScroll();

    return () => {
      window.removeEventListener("scroll", onScroll);
      if (lockTimerRef.current) {
        window.clearTimeout(lockTimerRef.current);
      }
    };
  }, [headings, offset, activeId]);

  return {
    activeId,
    scrollTo,
    setActiveId,
  };
};

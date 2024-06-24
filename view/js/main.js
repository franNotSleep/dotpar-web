const fileInput = document.getElementById("fileInput");
const inputTextarea = document.getElementById("inputTextarea");
const output = document.getElementById("output");
const compileBtn = document.getElementById("compileBtn");
const copyBtn = document.getElementById("copyBtn");
const toastInfo = document.getElementById("toastInfo");

import hljs from "https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/es/highlight.min.js";
import typescript from "https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.9.0/build/es/languages/typescript.min.js";
hljs.registerLanguage("typescript", typescript);
hljs.highlightAll();

fileInput.addEventListener("change", handleFileInputChange);
compileBtn.addEventListener("click", handleCompileClick);
copyBtn.addEventListener("click", handleCopyClick);

let compiledContent;

async function handleCopyClick() {
  if (!compiledContent) {
    return;
  }

  await navigator.clipboard.writeText(compiledContent);

  toastInfo.classList.remove("hidden");
  setTimeout(() => {
    toastInfo.classList.add("hidden");
  }, 5 * 1000);
}

async function handleCompileClick() {
  const response = await fetch("http://localhost:8080/compile", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ content: inputTextarea.value }),
  });

  const json = await response.json();
  compiledContent = json.compiled_content;

  const highlightedCode = hljs.highlight(json.compiled_content, {
    language: "typescript",
  }).value;

  output.innerHTML = highlightedCode;
}

async function handleFileInputChange() {
  const file = this.files[0];
  let content;

  if (file) {
    content = await file.text();
    inputTextarea.value = content;
  }
}

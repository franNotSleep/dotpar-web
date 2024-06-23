const fileInput = document.getElementById("fileInput");
const inputTextarea = document.getElementById("inputTextarea");
const output = document.getElementById("output");
const compileBtn = document.getElementById("compileBtn");

fileInput.addEventListener("change", handleFileInputChange);
compileBtn.addEventListener("click", handleCompileClick);

async function handleCompileClick() {
  const response = await fetch("http://localhost:8080/compile", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ content: inputTextarea.value }),
  });

  const json = await response.json();

  output.innerText = json.compiled_content;
}

async function handleFileInputChange() {
  const file = this.files[0];
  let content;

  if (file) {
    content = await file.text();
    inputTextarea.value = content;
  }
}

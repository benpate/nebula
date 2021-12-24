function makeEditor(quill) {

	quill.on("selection-change", function() {
		if (!quill.hasFocus()) {
			console.log(quill.getContents())
		}
	})
}
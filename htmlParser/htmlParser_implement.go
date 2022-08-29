package htmlParser;

type htmlStructure struct{
	rootPath string;
}

func New(rootPath string) HTMLParserInterface {
	return &htmlStructure{rootPath: rootPath};
}

func (*htmlStructure) Create(templateName string, data interface{}) (string, error) {
	return ""
}
package wip

func ReadInstalled(path string) (Packs, error) {
	folderNames, _ := listFolders(path)
	packs, _ := parseFolderNames(folderNames)

	return packs, nil
}

func parseFolderNames(folderNames []string) (Packs, error) {
	packs := CreatePacks()

	for _, folderName := range folderNames {
		pack, _ := CreatePackFromFolderName(folderName)
		packs = packs.Add(pack)
	}

	return packs, nil
}

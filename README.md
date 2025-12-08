| Command                                              | Explanation                     |
| ---------------------------------------------------- | ------------------------------- |
| `create <dict-name>`                                 | Create a new dictionary         |
| `delete <dict-name>`                                 | Delete existing dictionary      |
| `list`                                               | List existing dictionaries      |
| `list <dict-name>`                                   | List content of the dictionary  |
| `add <dict-name> -w <word> -d <definition> -t <tag>` | Create a word in a dictionary   |
| `remove <dict-name> <word>`                          | Remove a word from a dictionary |
| `search <dict-name> <word>`                          | Search for a word               |
| `search <dict-name> <tag> -t`                        | Search for words with tag       |

---

Run following command to dynamically generate documentation:

```bash
go run ./internal/tools/docgen -out ./docs -format markdown
```

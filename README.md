# **S3Flow**  
[![Go Version](https://img.shields.io/badge/Go-1.20+-blue)](https://golang.org/) [![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

S3Flow is a powerful tool designed to generate permutations of AWS S3 bucket names based on common prefixes, company names, and predefined environments. Whether you're testing cloud storage configurations or securing your buckets, S3Flow helps you craft comprehensive wordlists with ease.

---

## **Features**
- **Customizable Prefixes**: Generate bucket names using user-defined common prefixes.
- **Company-Specific Naming**: Incorporate your organization's name into bucket permutations.
- **Environment Variations**: Automatically include common environments like `dev`, `prod`, `stage`, etc.
- **Progress Bar**: Visualize the progress while saving generated wordlists to a file.
- **Efficient Deduplication**: Remove duplicate entries to ensure clean and concise output.

---

## **Installation**

### Prerequisites
- Go 1.20 or higher installed on your system.
- A text file (`common_bucket_prefixes.txt`) containing a list of prefixes to use in permutations.

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/S3Flow.git
   cd S3Flow
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the tool:
   ```bash
   go build -o s3flow main.go
   ```

4. (Optional) Move the binary to a directory in your PATH:
   ```bash
   sudo mv s3flow /usr/local/bin/
   ```

---

## **Usage**

S3Flow uses command-line flags to customize its behavior. Here's how to use it:

```bash
s3flow -w <common_prefix> -org <company_name> [-o <output_file>]
```

### Flags
| Flag       | Description                                   | Required? | Default Value         |
|------------|-----------------------------------------------|-----------|-----------------------|
| `-w`       | Common prefix for bucket names                | Yes       | N/A                   |
| `-org`     | Company or organization name                  | Yes       | N/A                   |
| `-o`       | Output file name                              | No        | `generated_wordlist.txt` |

### Example Usage
Generate a wordlist with the prefix `myapp` for the company `MyCompany`:
```bash
s3flow -w myapp -org MyCompany -o output.txt
```

This will:
1. Use `myapp` as the common prefix.
2. Incorporate `MyCompany` into the permutations.
3. Save the results to `output.txt` with a progress bar displayed during the save process.

---

## **How It Works**

S3Flow generates permutations by combining:
1. **Common Prefixes**: User-defined prefixes from the `common_bucket_prefixes.txt` file.
2. **Company Name**: The organization name provided via the `-org` flag.
3. **Environments**: Predefined environments like `dev`, `prod`, `stage`, etc.
4. **Formats**: Various naming conventions such as `prefix-word-env`, `prefix.word.env`, etc.

The tool ensures no duplicates are present in the final output.

---

## **Example Output**

Given the following inputs:
- Common Prefix: `myapp`
- Company Name: `MyCompany`
- Prefix Wordlist: `["data", "backup"]`

The generated wordlist might include:
```
myapp-data-dev
myapp-data-prod
myapp-backup-stage
myapp.MyCompany.dev
backup.myapp.production
...
```

---

## **Contributing**

We welcome contributions from the community! Here's how you can help:
1. Fork the repository.
2. Create a new branch for your feature or bug fix:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add feature or fix"
   ```
4. Push to your branch:
   ```bash
   git push origin feature-name
   ```
5. Open a pull request on GitHub.

---

## **License**

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.

---

## **Acknowledgments**

- Inspired by the need for efficient cloud storage bucket name generation.
- Thanks to the Go community for their amazing tools and libraries.

---

## **Contact**

For questions, suggestions, or feedback, feel free to reach out:
- Email: your.email@example.com
- GitHub: [@yourusername](https://github.com/yourusername)

---

This `README.md` provides a professional and user-friendly overview of your tool. It’s structured to make it easy for users to understand the purpose of S3Flow, how to install and use it, and how they can contribute. Let me know if you'd like to tweak anything further!
# **S3Flow**  
[![Go Version](https://img.shields.io/badge/Go-1.20+-blue)](https://golang.org/)

S3Flow is a utility tool designed to generate permutations of AWS S3 bucket names based on common prefixes, company names, and predefined environments. This tool is specially designed for Bug Bounty Hunters and Pentesters.

---

## **Installation**

### Prerequisites
- Go 1.20 or higher installed on your system.

### Steps

#### Using Go Install
```bash    
go install github.com/saeed0xf/s3flow@latest
```

#### Traditional way
1. Clone the repository:
   ```bash
   git clone https://github.com/saeed0xf/s3flow.git
   cd s3flow
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
s3flow -w <common_prefix_wordlist> -org <company_name> [-o <output_file>] [-medium] [-large] [-show-env] [-env-file <custom_env_file>]
```

Example Common Prefixes: [wordlist](https://raw.githubusercontent.com/koaj/aws-s3-bucket-wordlist/master/list.txt)

### Flags

| Flag       | Description                                   | Required? | Default Value         |
|------------|-----------------------------------------------|-----------|-----------------------|
| `-w`       | Common prefix wordlist file                   | Yes       | N/A                   |
| `-org`     | Company or organization name                  | Yes       | N/A                   |
| `-o`       | Output file name                              | No        | `generated_bucketlist.txt` |
| `-medium`  | Use medium-sized environment list             | No        | `false`               |
| `-large`   | Use large environment list                    | No        | `false`               |
| `-show-env`| Show all predefined environments              | No        | `false`               |
| `-env-file`| Path to a custom environment file             | No        | N/A                   |

### Environment Lists
- **Small**: Includes basic environments like `dev`, `prod`, `stage`, etc.
- **Medium**: Adds more environments like `uat`, `qa`, `sandbox`, etc.
- **Large**: Includes all possible environments, such as `backup`, `cdn`, `api`, etc.

### Example Usage

#### Generate a Small Wordlist (Default)
```bash
s3flow -w wordlist.txt -org MyCompany -o output.txt
```
This will:
1. Use `wordlist.txt` as the common prefixes.
2. Incorporate `MyCompany` into the permutations.
3. Save the results to `output.txt` using the small environment list.

#### Generate a Medium-Sized Wordlist
```bash
s3flow -w wordlist.txt -org MyCompany -o output.txt -medium
```
This will:
1. Use `wordlist.txt` as the common prefixes.
2. Incorporate `MyCompany` into the permutations.
3. Save the results to `output.txt` using the medium environment list.

#### Generate a Large Wordlist
```bash
s3flow -w wordlist.txt -org MyCompany -o output.txt -large
```
This will:
1. Use `wordlist.txt` as the common prefixes.
2. Incorporate `MyCompany` into the permutations.
3. Save the results to `output.txt` using the large environment list.

#### Show All Predefined Environments
```bash
s3flow -show-env
```
This will display all predefined environments (small, medium, and large) without generating a wordlist.

#### Use a Custom Environment File
```bash
s3flow -w wordlist.txt -org MyCompany -o output.txt -env-file custom_env.txt
```
This will:
1. Use `wordlist.txt` as the common prefixes.
2. Incorporate `MyCompany` into the permutations.
3. Use the environments provided in `custom_env.txt`.
4. Save the results to `output.txt`.

After generating the wordlist, you can use the wordlist with tools like [S3Scanner](https://github.com/sa7mon/S3Scanner).

---

## **How It Works**

S3Flow generates permutations by combining:
1. **Common Prefixes**: User-defined prefixes from the provided wordlist file.
2. **Company Name**: The organization name provided via the `-org` flag.
3. **Environments**: Predefined environments like `dev`, `prod`, `stage`, etc., selected based on the `-medium`, `-large`, or `-env-file` flags.
4. **Formats**: Various naming conventions such as `prefix-word-env`, `prefix.word.env`, etc.

The tool ensures no duplicates are present in the final output.

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

## **Contact**

For questions, suggestions, or feedback, feel free to reach out:
- Twitter/X: [Saeed0x1](https://x.com/saeed0x1)
- Linkedin: [Saeed0x1](https://www.linkedin.com/in/saeed0x1) 
- GitHub: [saeed0xf](https://github.com/saeed0xf)
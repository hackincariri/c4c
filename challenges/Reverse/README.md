# Reverse Engineering Challenge Submission

## What pattern should I follow?

- We do not accept any type of pre-compiled files.
- We will only accept challenges from compiled languages (Ex: C, Golang, C#).
- The source code must be in plain text.
- The compilation method must be provided.
- Any minification or obfuscation process must have step-by-step instructions provided.

## Example

Let's look at the example below, which is a reverse engineering challenge made with the C++ language. The code is in plain text without any obfuscation or minification.

1. Write the challenge code in plain text without any minification/obfuscation.

```C++
#include <iostream>
#include <string>

// dead code plus parameters in the file to display the flag
std::string unXORWithFF(const std::string& hexString) {
    std::string result;

    // Iterate through pairs of characters in the hex string
    for (size_t i = 0; i < hexString.length(); i += 2) {
        // Extract a pair of characters
        std::string pair = hexString.substr(i, 2);

        // Convert the pair from hex to decimal, XOR with FF, and convert back to ASCII
        char decryptedChar = static_cast<char>(std::stoi(pair, nullptr, 16) ^ 0xFF);

        // Append the decrypted character to the result string
        result.push_back(decryptedChar);
    }

    return result;
}

int main(int argc, char* argv[]) {
    std::string i;
    bool i2;
    std::cout << "Enter the password: ";
    std::cin >> i;

    {
        int x = 0;
        for (char c : i) {
            x += static_cast<int>(c);
        }

        if (x % 2 == 0) {
            i2 = true;
        } else {
            i2 = false;
        }
    }

    // More obfuscation: Random loops
    for (int i = 0; i < 10; ++i) {
        i2 = !i;
    }

    switch (i.length()) {
        case 5:
            i2 = !true;
            break;
        case 7:
            i2 = !i2;
            break;
        case 10:
            i2 = !false;
            break;
        default:
            break;
    }

    int c = 0;
    while (c < 5) {
        i2 = !i2;
        ++c;
    }

    if (strcmp(argv[1], "MZ") == 0 && strcmp(argv[2], "PE") == 0 && strcmp(argv[3], "HIK") == 0) {
        std::cout << unXORWithFF("B7B6B4A0B29E8D989E8D9A8BA0B79E9096938B9091A0C89ACA9BC9C9CFC6C6C6CACDC89CCD9CC699C8C8CAC9CA99CDC8CCCAC899CDC8") << std::endl;
    }

    return 0;
}
```

2. Explain how to execute the binary compilation.

```sh
g++ challenge.cpp -o challenge.exe
```

## How to submit a reverse engineering challenge

To submit your challenge, follow the steps below:

1. Create a private repository with your challenge.

2. Add the email organizacao@hackincariri.com.br as a contributor to the project.

3. Access the form at this [link](https://forms.gle/bnVjrsWELCpWpf1g8).

4. Put the link of the shared project.

5. Select the challenge type Reverse Engineering.

6. Put your nickname and your best email.

7. Wait for our contact. ;)



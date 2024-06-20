# Go Apps

```sh
go mod init phpguru.net/go-apps # tell go that this is a module
```

## Apps

### 0. Command-Line Flags

- https://gobyexample.com/command-line-flags

### 1. Leap Year

**Requirements**:

- Input: read a number from user
- Output: validate user's input value, and return the input year is leap year or not.

```
Please enter the year you wanna check :
The year [xxxx] is Leap Year and has 366 days
The year [xxxx] is not Leap Year and has 365 days
```

### 2. BMI

- A program calculate BMI (Body Mass Index) Metric. Based on BMI, it is possible to know whether that person is fat, thin or has an ideal weight.

> BMI = Weight/ [(Height)2]

In which, height is in meter and weight is in kilogram.

**Requirements**:

- Input: Read user's inputs via command line flags like this
- Output: bmi value and correspond bmi category

```sh
calculator -app=bmi -w=64 -h=1.64
```

> BMI Categories

```json
{
  "BMI_Categories": [
    {
      "Classification": "Low weight (thin)",
      "WHO": { "min": null, "max": 18.49 },
      "IDI_WPRO": { "min": null, "max": 18.49 }
    },
    {
      "Classification": "Normal",
      "WHO": { "min": 18.5, "max": 24.99 },
      "IDI_WPRO": { "min": 18.5, "max": 22.99 }
    },
    {
      "Classification": "Pre-obesity",
      "WHO": { "min": 25, "max": 29.99 },
      "IDI_WPRO": { "min": 23, "max": 24.99 }
    },
    {
      "Classification": "Obesity degree I",
      "WHO": { "min": 30, "max": 34.99 },
      "IDI_WPRO": { "min": 25, "max": 29.99 }
    },
    {
      "Classification": "Obesity class II",
      "WHO": { "min": 35, "max": 39.99 },
      "IDI_WPRO": { "min": 30, "max": null }
    },
    {
      "Classification": "Obesity degree III",
      "WHO": { "min": 40, "max": null },
      "IDI_WPRO": { "min": 40, "max": null }
    }
  ]
}
```

### Compound interest

![image](https://gist.github.com/assets/31009750/6ac3865f-8071-4ea3-8831-2a16fa8e2afc)

![image](https://gist.github.com/assets/31009750/0e435bc1-e428-4cd6-9cf6-4db8fa69b0dd)

Input:

```sh
# pv=1000000, r = 1%, n=12 months
calculator -app=ci -pv=1000000 -r=1 n=12
```

Output: AV (Future value) = ...

```math
A = P * (1+r)^n
```

## References

- [Command line flags](https://gobyexample.com/command-line-flags)

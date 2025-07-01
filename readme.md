# Classical Synchronization Problems in Go

This repository contains three separate Go applications that demonstrate solutions to well-known classical synchronization problems in computer science:

- **Dining Philosophers Problem**
- **Producer-Consumer Problem**
- **Sleeping Barber Problem**

These problems are fundamental in understanding concurrency, process synchronization, and resource sharing, making them popular examples in operating systems and concurrent programming courses.

---

## 📝 Problem Descriptions

### 1. Dining Philosophers Problem
A classic synchronization problem involving philosophers sitting around a table with forks placed between them. The challenge is to prevent deadlock and starvation while allowing all philosophers to eat.

### 2. Producer-Consumer Problem
Demonstrates synchronization between producer(s) generating data and consumer(s) processing data, ensuring that the shared buffer is accessed safely without race conditions or data loss.

### 3. Sleeping Barber Problem
Models a barbershop with a single barber and a limited number of waiting chairs. The problem illustrates proper synchronization to ensure customers are served efficiently without missing customers or causing deadlock.

---

## Concurrency Concepts Used

- **Goroutines**
- **Channels**
- **Mutexes**
- **WaitGroups**
- **Timeouts & Sleeping**
- **Basic inter-process synchronization patterns**
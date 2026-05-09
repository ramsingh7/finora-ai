# Finora AI

Finora AI is an AI-powered stock market analytics and prediction platform designed to help users analyze market behavior and generate intelligent forecasting insights. It combines time-series forecasting, technical analysis, sentiment analysis, and interactive dashboards to predict next-day stock prices, market trends, volatility, and buy/sell/hold signals.

## Backend Implementation Docs

- gRPC routing and method policy: `docs/BACKEND_GRPC_ROUTING.md`

## Overview

The goal of Finora AI is to build a portfolio-worthy and production-oriented financial intelligence platform that can evolve into a real-world SaaS product. The system uses historical stock data, technical indicators, financial news sentiment, and optional social media sentiment to provide smarter decision support for traders, learners, and analysts.

Finora AI is not positioned as a guaranteed trading system. It is an analytics and research platform that generates probabilistic insights based on available market data and machine learning models.

## Core Features

- Next-day stock price prediction
- Trend prediction (up/down/sideways)
- Volatility forecasting
- Buy / Sell / Hold signal generation
- Technical indicator analysis
- Candlestick chart visualization
- News sentiment analysis
- Optional social media sentiment integration
- Confidence scoring for predictions
- Portfolio and watchlist tracking
- Alert and notification system
- Backtesting and strategy simulation
- Explainable AI insights
- Model comparison dashboard
- Real-time or near-real-time analytics support

## Why Finora AI?

- Combines time-series forecasting with NLP-based sentiment analysis
- Uses both market data and external information signals
- Offers a professional dashboard and product-oriented architecture
- Strong portfolio project for AI, ML, Data Science, and Full Stack roles
- Scalable foundation for a real financial analytics startup

## Tech Stack

### Frontend
- React.js / Next.js
- Tailwind CSS
- Charting libraries such as Recharts or ApexCharts

### Backend
- FastAPI or Node.js
- REST APIs
- WebSocket support for live updates

### AI / ML
- Python
- Pandas
- NumPy
- Scikit-learn
- XGBoost / LightGBM
- LSTM / GRU
- Transformer-based time-series models
- NLP sentiment models such as FinBERT

### Database
- PostgreSQL for structured data
- MongoDB for flexible document-based sentiment/news storage
- Redis for caching and fast access

### DevOps / Deployment
- Docker
- CI/CD pipelines
- Cloud deployment
- Scheduled jobs for retraining and data ingestion

## Problem Statement

Stock markets are highly dynamic, noisy, and influenced by multiple factors such as price history, technical patterns, news events, macroeconomic conditions, and investor sentiment. Traditional analysis methods often fail to combine these diverse signals effectively.

Finora AI addresses this by building an integrated prediction and analytics platform that can:

- analyze historical market behavior
- extract useful technical features
- understand financial sentiment from text data
- forecast potential price movement
- generate research-oriented buy/sell/hold signals
- provide visual explanations and backtesting insights

## Prediction Targets

Finora AI can support multiple prediction tasks:

- Next-day closing price prediction
- Next-day return prediction
- Trend direction classification
- Volatility prediction
- Buy / Sell / Hold signal classification
- Confidence score generation
- Multi-day forecasting
- Risk scoring and anomaly detection

## Data Sources

Possible data sources include:

- Historical OHLCV stock data
- Intraday market data
- Market index data
- Sector-wise stock data
- Fundamental company data
- Earnings data
- Financial news articles
- News headlines
- Social media sentiment data
- Macroeconomic indicators
- Global market sentiment signals

## Machine Learning Approach

Finora AI supports multiple modeling approaches depending on project stage and complexity.

### Traditional ML Models
- Linear Regression
- Logistic Regression
- Random Forest
- XGBoost
- LightGBM
- SVM

### Deep Learning Models
- LSTM
- GRU
- CNN-LSTM
- Transformer-based forecasting
- Temporal Fusion Transformer

### NLP Models
- Rule-based sentiment scoring
- TF-IDF + classical ML
- FinBERT / finance-specific transformer models
- Named entity extraction for stock/company mapping

### Hybrid Models
The most powerful version of Finora AI combines:
- historical time-series features
- technical indicators
- sentiment signals
- market index signals
- ensemble predictions

## High-Level Architecture

```text
User Interface (React / Next.js)
        |
        v
Backend API Layer (FastAPI / Node.js)
        |
        +----------------------+
        |                      |
        v                      v
Prediction Service        Data & Sentiment Service
(ML / DL Models)          (Market + News + Social Data)
        |                      |
        +-----------+----------+
                    |
                    v
         PostgreSQL / MongoDB / Redis
                    |
                    v
         Training Pipeline / Scheduler / Logs
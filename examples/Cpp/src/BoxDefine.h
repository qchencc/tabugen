﻿// This file is auto-generated by Tabugen v0.10.0, DO NOT EDIT!

#pragma once

#include <stdint.h>
#include <string>
#include <vector>
#include <unordered_map>
#include <functional>


namespace config {

// 
struct BoxProbabilityDefine 
{
    struct ProbabilityGoodsDefine 
    {
        std::string  GoodsID;               // 物品id
        uint32_t     Num = 0;               // 物品数量
        uint32_t     Probability = 0;       // 物品概率
    };

    std::string                          ID;                    // ID
    int                                  Total = 0;             // 奖励总数
    int                                  Time = 0;              // 冷却时间
    bool                                 Repeat = false;        // 是否可重复
    std::vector<ProbabilityGoodsDefine>  ProbabilityGoods;    // 

    static int ParseFrom(const std::unordered_map<std::string, std::string>& fields, BoxProbabilityDefine* ptr);
};

} // namespace config

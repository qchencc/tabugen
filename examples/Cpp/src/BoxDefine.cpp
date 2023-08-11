﻿// This file is auto-generated by Tabular v0.10.0, DO NOT EDIT!

#include "BoxDefine.h"
#include <stddef.h>
#include "Conv.h"

using namespace std;

#ifndef ASSERT
#define ASSERT assert
#endif


namespace config {

// parse BoxProbabilityDefine from string fields
int BoxProbabilityDefine::ParseFrom(const unordered_map<string, string>& record, BoxProbabilityDefine* ptr)
{
    ASSERT(ptr != nullptr);
    ptr->ID = parseField<std::string>(record, "ID");
    ptr->Total = parseField<int>(record, "Total");
    ptr->Time = parseField<int>(record, "Time");
    ptr->Repeat = parseField<bool>(record, "Repeat");
    for (size_t i = 0; i < record.size(); i++)
    {
        BoxProbabilityDefine::ProbabilityGoodsDefine val;
        {
            auto key = stringPrintf("GoodsID[%d]", i);
            auto iter = record.find(key);
            if (iter == record.end()) {
                break;
            }
            val.GoodsID = parseField<std::string>(record, key);
        }
        {
            auto key = stringPrintf("Num[%d]", i);
            val.Num = parseField<uint32_t>(record, key);
        }
        {
            auto key = stringPrintf("Probability[%d]", i);
            val.Probability = parseField<uint32_t>(record, key);
        }
        ptr->ProbabilityGoods.push_back(val);
    }
    ptr->ProbabilityGoods.shrink_to_fit();
    return 0;
}


} // namespace config 

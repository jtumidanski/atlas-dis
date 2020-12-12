package com.atlas.dis.model;

public record MonsterDropData(int id, int monsterId, int itemId, int maximumQuantity, int minimumQuantity, int chance) {
}
